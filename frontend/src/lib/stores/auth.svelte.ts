/**
 * Authentication Store (Svelte 5 Runes)
 *
 * Holds client-side auth state.  The source of truth remains the server
 * (HttpOnly cookies + server-side load functions).  This store is seeded
 * from page data and kept in sync as the user navigates.
 */

import type { User } from '$lib/types/auth';

interface AuthState {
	user: User | null;
	isAuthenticated: boolean;
	isLoading: boolean;
	error: string | null;
}

const authState = $state<AuthState>({
	user: null,
	isAuthenticated: false,
	isLoading: false,
	error: null
});

class AuthStore {
	// ── Accessors ──────────────────────────────────────────────────────────

	get user(): User | null {
		return authState.user;
	}

	get isAuthenticated(): boolean {
		return authState.isAuthenticated;
	}

	get isLoading(): boolean {
		return authState.isLoading;
	}

	get error(): string | null {
		return authState.error;
	}

	// ── Mutations ──────────────────────────────────────────────────────────

	/**
	 * Called from the root layout after the server load resolves so that
	 * client-side components always have up-to-date user information.
	 */
	setUser(user: User | null) {
		authState.user = user;
		authState.isAuthenticated = user !== null;
		authState.error = null;
	}

	setLoading(loading: boolean) {
		authState.isLoading = loading;
	}

	setError(error: string | null) {
		authState.error = error;
	}

	clearError() {
		authState.error = null;
	}

	/**
	 * Clear local state on logout.  The actual session invalidation is done
	 * server-side via the /logout form action.
	 */
	clear() {
		authState.user = null;
		authState.isAuthenticated = false;
		authState.error = null;
		authState.isLoading = false;
	}
}

export const authStore = new AuthStore();
