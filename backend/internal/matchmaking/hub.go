package matchmaking

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

// PlayerInQueue holds a client and what they're queueing for
type PlayerInQueue struct {
	Client  *Client
	GameIDs []uint
}

// ClientMessage is a wrapper for passing messages from
// readPump to the hub's Run loop.
type ClientMessage struct {
	client  *Client
	message *Message
}

// Hub maintains the set of active clients and handles connection logic.
type Hub struct {
	register   chan *Client
	unregister chan *Client

	// This map stores the *active* client for a given userID.
	userClients map[uint]*Client
	allClients  map[*Client]bool
	incoming    chan *ClientMessage

	clientsMutex sync.RWMutex

	queue      map[uint][]*PlayerInQueue
	queueMutex sync.Mutex

	// DB is used for side-effects like incrementing match counts
	DB *gorm.DB
}

func NewHub(db *gorm.DB) *Hub {
	return &Hub{
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		userClients: make(map[uint]*Client),
		allClients:  make(map[*Client]bool),
		incoming:    make(chan *ClientMessage),
		queue:       make(map[uint][]*PlayerInQueue),
		DB:          db,
	}
}

// GetOnlineUserIDs returns a set of currently connected user IDs.
func (h *Hub) GetOnlineUserIDs() map[uint]struct{} {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()
	set := make(map[uint]struct{}, len(h.userClients))
	for id := range h.userClients {
		set[id] = struct{}{}
	}
	return set
}

// Helper to create a new S2C (Server to Client) message
func newS2CMessage(msgType string, payload interface{}) []byte {
	payloadBytes, _ := json.Marshal(payload)
	msg, _ := json.Marshal(&Message{
		Type:    msgType,
		Payload: payloadBytes,
	})
	return msg
}

// broadcastToAll sends a message to all connected clients
func (h *Hub) broadcastToAll(message []byte) {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()
	for client := range h.allClients {
		// Use a select to prevent blocking on a slow client
		select {
		case client.send <- message:
		default:
			// Client's send buffer is full, they might be lagging
		}
	}
}

// broadcastQueueUpdate calculates queue size and sends it to all
func (h *Hub) broadcastQueueUpdate() {
	h.queueMutex.Lock()
	size := 0
	for _, q := range h.queue {
		size += len(q)
	}
	h.queueMutex.Unlock()

	msg := newS2CMessage("queue_update", &QueueUpdatePayload{
		QueueSize: size,
	})
	h.broadcastToAll(msg)
}

// 💡 NEW: The Matchmaking Logic
func (h *Hub) tryFindMatch() {
	h.queueMutex.Lock()
	defer h.queueMutex.Unlock()

	// Loop over all game queues
	for gameID, players := range h.queue {
		// ⭐️ YOUR LOGIC ⭐️
		if len(players) >= 5 {
			log.Printf("Found match for game %d!", gameID)

			// 1. Get the 5 players
			matchedPlayers := players[:5]
			// 2. Remove them from the queue
			h.queue[gameID] = players[5:]

			// 3. Create the room and send them the message
			// We do this in a goroutine so we don't block the lock
			go func(matched []*PlayerInQueue) {
				// In a real app, you'd generate a unique URL
				roomURL := "/ws/room/abc12345"
				msg := newS2CMessage("match_found", &MatchFoundPayload{
					RoomURL: roomURL,
				})

				for _, p := range matched {
					p.Client.send <- msg
					h.unregister <- p.Client

					// Increment total_matches for each matched player (best-effort)
					if h.DB != nil {
						h.DB.Exec(
							"UPDATE profiles SET total_matches = total_matches + 1 WHERE user_id = ?",
							p.Client.userID,
						)
					}
				}

				// Tell everyone else the queue size changed
				h.broadcastQueueUpdate()

			}(matchedPlayers)
		}
	}
}

// closeClient is a helper to safely close a client's
// 'send' channel and connection.
func (h *Hub) closeClient(c *Client, message string) {
	c.closeOnce.Do(func() {
		log.Printf("Closing connection for user %d: %s", c.userID, message)
		// We close the 'send' channel to signal to writePump to stop.
		close(c.send)
		// We close the websocket connection.
		c.conn.Close()
	})
}

func (h *Hub) Run() {
	// 💡 NEW: Start a ticker to run the match logic
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 💡 Run matching logic periodically
			h.tryFindMatch()

		case client := <-h.register:
			h.clientsMutex.Lock()
			if existingClient, ok := h.userClients[client.userID]; ok {
				h.closeClient(existingClient, "new connection replaced this one")
				delete(h.allClients, existingClient) // 👈 Don't forget
			}
			h.userClients[client.userID] = client
			h.allClients[client] = true // 👈 NEW
			h.clientsMutex.Unlock()

			log.Printf("Client registered for user %d", client.userID)
			// Send the current queue size to the new client
			h.broadcastQueueUpdate()

		case client := <-h.unregister:
			h.clientsMutex.Lock()
			if existingClient, ok := h.userClients[client.userID]; ok && existingClient == client {
				delete(h.userClients, client.userID)
				delete(h.allClients, client) // 👈 NEW
				h.closeClient(client, "client disconnected")
				log.Printf("Client unregistered for user %d", client.userID)

				// 💡 NEW: Check if they were in the queue
				h.removePlayerFromAllQueues(client)

			} else {
				h.closeClient(client, "stale disconnect")
			}
			h.clientsMutex.Unlock()

		case clientMsg := <-h.incoming:
			// 💡 NEW: Handle incoming messages
			h.handleClientMessage(clientMsg)
		}
	}
}

// 💡 NEW: Message handler
func (h *Hub) handleClientMessage(cm *ClientMessage) {
	switch cm.message.Type {
	case "join_queue":
		var payload JoinQueuePayload
		if err := json.Unmarshal(cm.message.Payload, &payload); err != nil {
			log.Printf("Failed to parse join_queue: %v", err)
			return
		}

		// Simplified: We just use the *first* gameID
		if len(payload.GameIDs) == 0 {
			return
		}
		gameID := payload.GameIDs[0]

		log.Printf("Client %d joining queue for game %d", cm.client.userID, gameID)

		player := &PlayerInQueue{
			Client:  cm.client,
			GameIDs: payload.GameIDs,
		}

		h.queueMutex.Lock()
		h.queue[gameID] = append(h.queue[gameID], player)
		h.queueMutex.Unlock()

		// Tell *everyone* the queue size has changed
		h.broadcastQueueUpdate()

	case "leave_queue":
		log.Printf("Client %d leaving queue", cm.client.userID)
		// We already built the logic to handle this!
		// It's safe to call even if they aren't in a queue.
		h.removePlayerFromAllQueues(cm.client)
		// broadcastQueueUpdate is already called inside removePlayerFromAllQueues
	}
}

// 💡 NEW: Helper to clean up queues on disconnect
func (h *Hub) removePlayerFromAllQueues(client *Client) {
	h.queueMutex.Lock()
	defer h.queueMutex.Unlock()

	needsBroadcast := false
	for gameID, players := range h.queue {
		// We create a new slice, filtering out the client
		newQueue := make([]*PlayerInQueue, 0, len(players))
		for _, p := range players {
			if p.Client != client {
				newQueue = append(newQueue, p)
			} else {
				needsBroadcast = true
			}
		}
		h.queue[gameID] = newQueue
	}

	if needsBroadcast {
		go h.broadcastQueueUpdate()
	}
}
