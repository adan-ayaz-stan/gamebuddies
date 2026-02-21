import { api, forwardSetCookies } from '$lib/server/api';
import type { Handle, HandleServerError } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';

/**
 * Authentication hook – runs on every server request.
 * Reads HttpOnly cookies set by the Go backend, validates them by calling
 * the backend profile endpoint, and populates event.locals.user.
 * Handles silent token refresh when the access token has expired.
 */
const authenticationHandle: Handle = async ({ event, resolve }) => {
	const accessToken = event.cookies.get('gm_access_token');
	const refreshToken = event.cookies.get('gm_refresh_token');

	// ── 1. Try access token ────────────────────────────────────────────────
	if (accessToken) {
		const user = await api.auth.getProfile(accessToken);
		if (user) {
			event.locals.user = user;
			return resolve(event);
		}
	}

	// ── 2. Access token missing or expired – try silent refresh ───────────
	if (refreshToken) {
		const refreshResp = await api.auth.refresh(refreshToken);

		if (refreshResp.ok) {
			// Forward new cookies the backend issued back to the browser
			forwardSetCookies(refreshResp, event.cookies);

			// Fetch profile with the freshly issued access token
			const newAccessToken = event.cookies.get('gm_access_token');
			if (newAccessToken) {
				const user = await api.auth.getProfile(newAccessToken);
				if (user) {
					event.locals.user = user;
					return resolve(event);
				}
			}
		} else {
			// Refresh token is invalid / expired – clear stale cookies
			event.cookies.delete('gm_access_token', { path: '/' });
			event.cookies.delete('gm_refresh_token', { path: '/' });
		}
	}

	// ── 3. Unauthenticated ─────────────────────────────────────────────────
	event.locals.user = null;
	return resolve(event);
};

/**
 * Security headers hook.
 */
const securityHandle: Handle = async ({ event, resolve }) => {
	const response = await resolve(event);
	response.headers.set('X-Frame-Options', 'SAMEORIGIN');
	response.headers.set('X-Content-Type-Options', 'nosniff');
	response.headers.set('Referrer-Policy', 'strict-origin-when-cross-origin');
	return response;
};

export const handle = sequence(authenticationHandle, securityHandle);

export const handleError: HandleServerError = async ({ error, event, status, message }) => {
	const errorId = crypto.randomUUID();
	console.error(`[${errorId}] ${status}:`, error, { url: event.url.pathname });
	return {
		message: status === 500 ? 'An unexpected error occurred' : message,
		errorId
	};
};
