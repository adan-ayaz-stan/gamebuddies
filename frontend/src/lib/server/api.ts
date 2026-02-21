import type {
	Friend,
	LoginRequest,
	LoginResponse,
	RegisterRequest,
	RegisterResponse,
	Stats,
	User
} from '$lib/types/auth';
import type { Cookies } from '@sveltejs/kit';

export const BACKEND_URL = 'http://localhost:8080/v1';

// ---------------------------------------------------------------------------
// Cookie forwarding helper
// Reads Set-Cookie headers from a backend response and mirrors them via
// SvelteKit's cookies API so the browser receives them transparently.
// ---------------------------------------------------------------------------
export function forwardSetCookies(response: Response, cookies: Cookies) {
	const raw: string[] =
		// Node 18+ exposes getSetCookie(); fall back to splitting the header value.
		typeof (response.headers as any).getSetCookie === 'function'
			? (response.headers as any).getSetCookie()
			: (response.headers.get('set-cookie') ?? '').split(/,(?=[^ ])/).filter(Boolean);

	for (const header of raw) {
		const parts = header.split(';').map((s) => s.trim());
		const [nameVal, ...attrs] = parts;
		const eqIdx = nameVal.indexOf('=');
		if (eqIdx === -1) continue;

		const name = nameVal.slice(0, eqIdx).trim();
		const value = nameVal.slice(eqIdx + 1).trim();

		if (!name) continue;

		let maxAge: number | undefined;
		let path = '/';
		for (const attr of attrs) {
			const [k, v] = attr.split('=');
			const key = k.trim().toLowerCase();
			if (key === 'max-age') maxAge = parseInt(v ?? '0', 10);
			if (key === 'path') path = v?.trim() ?? '/';
		}

		cookies.set(name, value, {
			path,
			httpOnly: true,
			sameSite: 'lax',
			secure: false, // false for local dev; production should use true
			...(maxAge !== undefined ? { maxAge } : {})
		});
	}
}

// ---------------------------------------------------------------------------
// Auth API calls – all cookie-based; tokens are never kept in JS memory
// ---------------------------------------------------------------------------
export const api = {
	auth: {
		/** POST /auth/sign-in – returns body and sets HttpOnly cookies */
		login: async (credentials: LoginRequest): Promise<Response> => {
			return fetch(`${BACKEND_URL}/auth/sign-in`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(credentials)
			});
		},

		/** POST /auth/sign-up */
		register: async (credentials: RegisterRequest): Promise<Response> => {
			return fetch(`${BACKEND_URL}/auth/sign-up`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(credentials)
			});
		},

		/** POST /auth/refresh – forwards the refresh cookie; returns response w/ new cookies */
		refresh: async (refreshToken: string): Promise<Response> => {
			return fetch(`${BACKEND_URL}/auth/refresh`, {
				method: 'POST',
				headers: { Cookie: `gm_refresh_token=${refreshToken}` }
			});
		},

		/** POST /auth/logout – forwards both cookies to let backend revoke the token */
		logout: async (accessToken: string, refreshToken: string): Promise<Response> => {
			const cookieParts: string[] = [];
			if (accessToken) cookieParts.push(`gm_access_token=${accessToken}`);
			if (refreshToken) cookieParts.push(`gm_refresh_token=${refreshToken}`);
			return fetch(`${BACKEND_URL}/auth/logout`, {
				method: 'POST',
				headers: { Cookie: cookieParts.join('; ') }
			});
		},

		/** GET /profile – forwards the access cookie; returns { user_id, username, display_name, avatar_url } */
		getProfile: async (accessToken: string): Promise<User | null> => {
			const resp = await fetch(`${BACKEND_URL}/profile`, {
				method: 'GET',
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
			if (!resp.ok) return null;
			const data = await resp.json();
			return {
				id: data.user_id,
				username: data.username,
				display_name: data.display_name,
				avatar_url: data.avatar_url || undefined
			} satisfies User;
		}
	},

	// -------------------------------------------------------------------------
	// User / Social API calls
	// -------------------------------------------------------------------------
	users: {
		/** GET /friends – returns list of accepted friends with online status */
		getFriends: async (accessToken: string): Promise<Friend[]> => {
			const resp = await fetch(`${BACKEND_URL}/friends`, {
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
			if (!resp.ok) return [];
			const data = await resp.json();
			return Array.isArray(data) ? data : [];
		},

		/** POST /friends – sends a friend request (pending) */
		sendFriendRequest: async (accessToken: string, username: string): Promise<Response> => {
			return fetch(`${BACKEND_URL}/friends`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Cookie: `gm_access_token=${accessToken}`
				},
				body: JSON.stringify({ username })
			});
		},

		/** GET /friends/requests – returns incoming pending friend requests */
		getFriendRequests: async (
			accessToken: string
		): Promise<import('$lib/types/auth').FriendRequest[]> => {
			const resp = await fetch(`${BACKEND_URL}/friends/requests`, {
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
			if (!resp.ok) return [];
			const data = await resp.json();
			return Array.isArray(data) ? data : [];
		},

		/** POST /friends/requests/:id/accept */
		acceptFriendRequest: async (accessToken: string, requestId: number): Promise<Response> => {
			return fetch(`${BACKEND_URL}/friends/requests/${requestId}/accept`, {
				method: 'POST',
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
		},

		/** POST /friends/requests/:id/decline */
		declineFriendRequest: async (accessToken: string, requestId: number): Promise<Response> => {
			return fetch(`${BACKEND_URL}/friends/requests/${requestId}/decline`, {
				method: 'POST',
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
		},

		/** DELETE /friends/:friendId – remove a friend */
		removeFriend: async (accessToken: string, friendId: number): Promise<Response> => {
			return fetch(`${BACKEND_URL}/friends/${friendId}`, {
				method: 'DELETE',
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
		},

		/** GET /stats – returns session / match stats for the authenticated user */
		getStats: async (accessToken: string): Promise<Stats | null> => {
			const resp = await fetch(`${BACKEND_URL}/stats`, {
				headers: { Cookie: `gm_access_token=${accessToken}` }
			});
			if (!resp.ok) return null;
			return resp.json() as Promise<Stats>;
		},

		/** PATCH /profile – update display_name */
		updateProfile: async (
			accessToken: string,
			data: { display_name?: string }
		): Promise<Response> => {
			return fetch(`${BACKEND_URL}/profile`, {
				method: 'PATCH',
				headers: {
					'Content-Type': 'application/json',
					Cookie: `gm_access_token=${accessToken}`
				},
				body: JSON.stringify(data)
			});
		}
	}
};

/** Parse the login response body into a typed object */
export async function parseLoginResponse(response: Response): Promise<LoginResponse> {
	return response.json() as Promise<LoginResponse>;
}

/** Parse the register response body into a typed object */
export async function parseRegisterResponse(response: Response): Promise<RegisterResponse> {
	return response.json() as Promise<RegisterResponse>;
}
