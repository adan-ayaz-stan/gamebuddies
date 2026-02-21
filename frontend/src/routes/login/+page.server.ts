import { api, forwardSetCookies } from '$lib/server/api';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, url }) => {
	// Already authenticated → redirect away
	if (locals.user) {
		const redirectTo = url.searchParams.get('redirectTo') ?? '/dashboard';
		redirect(302, redirectTo);
	}
	return {
		registered: url.searchParams.get('registered') === '1'
	};
};

export const actions: Actions = {
	default: async ({ request, cookies, url }) => {
		const form = await request.formData();
		const username = form.get('username')?.toString().trim() ?? '';
		const password = form.get('password')?.toString() ?? '';

		if (!username || !password) {
			return fail(400, { username, error: 'Username and password are required.' });
		}

		let response: Response;
		try {
			response = await api.auth.login({ username, password });
		} catch {
			return fail(503, { username, error: 'Could not reach the server. Try again later.' });
		}

		if (!response.ok) {
			let message = 'Invalid credentials.';
			try {
				const data = await response.json();
				message = data.message || data.error || message;
			} catch {
				/* ignore */
			}
			return fail(response.status, { username, error: message });
		}

		// Mirror the HttpOnly cookies the backend set onto the browser response
		forwardSetCookies(response, cookies);

		const redirectTo = url.searchParams.get('redirectTo') ?? '/dashboard';
		redirect(302, redirectTo);
	}
};
