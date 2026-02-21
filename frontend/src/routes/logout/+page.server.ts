import { api } from '$lib/server/api';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	default: async ({ cookies }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const refreshToken = cookies.get('gm_refresh_token') ?? '';

		// Revoke tokens on the backend (best-effort)
		try {
			if (accessToken || refreshToken) {
				await api.auth.logout(accessToken, refreshToken);
			}
		} catch {
			/* Backend may be unreachable – still clear local cookies */
		}

		// Clear auth cookies regardless of backend response
		cookies.delete('gm_access_token', { path: '/' });
		cookies.delete('gm_refresh_token', { path: '/' });

		redirect(302, '/login');
	}
};
