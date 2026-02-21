package matchmaking

import "encoding/json"

// Message defines the standard C2S (Client to Server)
// and S2C (Server to Client) websocket message.
type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"` // Flexible data
}

// === Client to Server (C2S) Payloads ===

// JoinQueuePayload is the data C2S when they want to queue.
type JoinQueuePayload struct {
	GameIDs []uint `json:"gameIDs"`
	// We could add other things here later, like:
	// PlayerMMR int `json:"playerMMR"` 
}

// === Server to Client (S2C) Payloads ===

// QueueUpdatePayload is the data S2C to all clients.
type QueueUpdatePayload struct {
	QueueSize int `json:"queueSize"`
}

// MatchFoundPayload is the data S2C *only* to matched players.
type MatchFoundPayload struct {
	RoomURL string `json:"roomURL"`
}