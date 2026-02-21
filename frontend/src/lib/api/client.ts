/**
 * GameBuddies API Client
 * 
 * Clean slate - Build your API integration from here!
 * 
 * Base URL: http://localhost:8080
 * 
 */

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

class ApiClient {
	/**
	 * Base request method - handles all HTTP requests
	 * Add authentication, error handling, etc. here
	 */
	private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
		const url = `${API_BASE_URL}${endpoint}`;

		const headers: Record<string, string> = {
			'Content-Type': 'application/json',
			...(options.headers as Record<string, string>),
		};

		const response = await fetch(url, {
			...options,
			headers,
		});

		if (!response.ok) {
			const error = await response.json().catch(() => ({ error: 'Request failed' }));
			throw new Error(error.error || `HTTP ${response.status}`);
		}

		return response.json();
	}

	// ==========================================
	// START HERE: Add your API methods below
	// ==========================================
}

export const apiClient = new ApiClient();

