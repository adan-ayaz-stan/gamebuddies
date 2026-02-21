import { browser } from '$app/environment';

interface WebSocketOptions {
	url: string;
	onMessage?: (data: unknown) => void;
	onConnect?: () => void;
	onDisconnect?: () => void;
	onError?: (error: Event) => void;
	autoReconnect?: boolean;
	reconnectInterval?: number;
}

export class WebSocketManager {
	private ws: WebSocket | null = null;
	private options: WebSocketOptions;
	private reconnectTimer: ReturnType<typeof setTimeout> | null = null;
	private shouldReconnect = true;
	private isConnected = $state(false);

	constructor(options: WebSocketOptions) {
		this.options = {
			autoReconnect: true,
			reconnectInterval: 3000,
			...options
		};

		if (browser) {
			this.connect();
		}
	}

	get connected() {
		return this.isConnected;
	}

	private connect() {
		try {
			this.ws = new WebSocket(this.options.url);

			this.ws.onopen = () => {
				this.isConnected = true;
				this.options.onConnect?.();
				// Clear any reconnection timer
				if (this.reconnectTimer) {
					clearTimeout(this.reconnectTimer);
					this.reconnectTimer = null;
				}
			};

			this.ws.onmessage = (event) => {
				try {
					const data = JSON.parse(event.data);
					this.options.onMessage?.(data);
				} catch (error) {
					console.error('Failed to parse WebSocket message:', error);
				}
			};

			this.ws.onclose = () => {
				this.isConnected = false;
				this.options.onDisconnect?.();
				this.attemptReconnect();
			};

			this.ws.onerror = (error) => {
				console.error('WebSocket error:', error);
				this.options.onError?.(error);
			};
		} catch (error) {
			console.error('Failed to create WebSocket connection:', error);
			this.attemptReconnect();
		}
	}

	private attemptReconnect() {
		if (
			this.shouldReconnect &&
			this.options.autoReconnect &&
			!this.reconnectTimer
		) {
			this.reconnectTimer = setTimeout(() => {
				this.reconnectTimer = null;
				this.connect();
			}, this.options.reconnectInterval);
		}
	}

	send(data: unknown) {
		if (this.ws?.readyState === WebSocket.OPEN) {
			this.ws.send(JSON.stringify(data));
		} else {
			console.warn('WebSocket is not connected');
		}
	}

	disconnect() {
		this.shouldReconnect = false;
		if (this.reconnectTimer) {
			clearTimeout(this.reconnectTimer);
			this.reconnectTimer = null;
		}
		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}
		this.isConnected = false;
	}
}

export function createWebSocket(options: WebSocketOptions) {
	return new WebSocketManager(options);
}