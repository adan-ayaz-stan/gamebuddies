/**
 * Matchmaking Store (Svelte 5 Runes)
 *
 * Manages WebSocket connection to the matchmaking hub, queue state,
 * and match notifications.
 *
 * WS message protocol (C2S):
 *   { type: "join_queue",  payload: { gameIDs: number[] } }
 *   { type: "leave_queue", payload: null }
 *
 * WS message protocol (S2C):
 *   { type: "queue_update", payload: { queueSize: number } }
 *   { type: "match_found",  payload: { roomURL: string } }
 */

const WS_URL = `ws://localhost:8080/v1/ws/matchmaking`;

interface WsMessage {
	type: string;
	payload: unknown;
}

interface MatchmakingState {
	connected: boolean;
	inQueue: boolean;
	queueSize: number;
	matchFound: boolean;
	roomURL: string | null;
	error: string | null;
}

const state = $state<MatchmakingState>({
	connected: false,
	inQueue: false,
	queueSize: 0,
	matchFound: false,
	roomURL: null,
	error: null
});

class MatchmakingStore {
	private ws: WebSocket | null = null;
	private reconnectTimer: ReturnType<typeof setTimeout> | null = null;
	private shouldReconnect = false;

	// ── Accessors ────────────────────────────────────────────────────────

	get connected() {
		return state.connected;
	}
	get inQueue() {
		return state.inQueue;
	}
	get queueSize() {
		return state.queueSize;
	}
	get matchFound() {
		return state.matchFound;
	}
	get roomURL() {
		return state.roomURL;
	}
	get error() {
		return state.error;
	}

	// ── Connection management ────────────────────────────────────────────

	connect() {
		if (
			this.ws &&
			(this.ws.readyState === WebSocket.OPEN || this.ws.readyState === WebSocket.CONNECTING)
		) {
			return;
		}
		this.shouldReconnect = true;
		this._openSocket();
	}

	disconnect() {
		this.shouldReconnect = false;
		if (this.reconnectTimer) clearTimeout(this.reconnectTimer);
		if (this.ws) {
			this.ws.close(1000, 'user disconnected');
			this.ws = null;
		}
		state.connected = false;
		state.inQueue = false;
		state.matchFound = false;
		state.roomURL = null;
		state.error = null;
	}

	private _openSocket() {
		try {
			this.ws = new WebSocket(WS_URL);
		} catch {
			state.error = 'Failed to create WebSocket';
			return;
		}

		this.ws.onopen = () => {
			state.connected = true;
			state.error = null;
			if (this.reconnectTimer) {
				clearTimeout(this.reconnectTimer);
				this.reconnectTimer = null;
			}
		};

		this.ws.onclose = (evt) => {
			state.connected = false;
			state.inQueue = false;
			if (!evt.wasClean && this.shouldReconnect) {
				this.reconnectTimer = setTimeout(() => this._openSocket(), 3000);
			}
		};

		this.ws.onerror = () => {
			state.error = 'WebSocket connection error';
		};

		this.ws.onmessage = (evt) => {
			try {
				const msg = JSON.parse(evt.data as string) as WsMessage;
				this._handleMessage(msg);
			} catch {
				console.warn('[matchmaking] could not parse message', evt.data);
			}
		};
	}

	private _handleMessage(msg: WsMessage) {
		switch (msg.type) {
			case 'queue_update': {
				const p = msg.payload as { queueSize: number };
				state.queueSize = p.queueSize ?? 0;
				break;
			}
			case 'match_found': {
				const p = msg.payload as { roomURL: string };
				state.matchFound = true;
				state.roomURL = p.roomURL ?? null;
				state.inQueue = false;
				break;
			}
			default:
				console.warn('[matchmaking] unknown message type:', msg.type);
		}
	}

	// ── Queue actions ────────────────────────────────────────────────────

	joinQueue(gameIDs: number[]) {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			state.error = 'Not connected to matchmaking server';
			return;
		}
		this.ws.send(JSON.stringify({ type: 'join_queue', payload: { gameIDs } }));
		state.inQueue = true;
		state.error = null;
	}

	leaveQueue() {
		if (this.ws && this.ws.readyState === WebSocket.OPEN) {
			this.ws.send(JSON.stringify({ type: 'leave_queue', payload: null }));
		}
		state.inQueue = false;
	}

	clearMatch() {
		state.matchFound = false;
		state.roomURL = null;
	}
}

export const matchmakingStore = new MatchmakingStore();
