import { api } from '$lib/server/api';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.user) redirect(302, '/dashboard');
	return {};
};

export const actions: Actions = {
	default: async ({ request }) => {
		const form = await request.formData();
		const display_name = form.get('display_name')?.toString().trim() ?? '';
		const username = form.get('username')?.toString().trim().toLowerCase() ?? '';
		const password = form.get('password')?.toString() ?? '';
		const confirm_password = form.get('confirm_password')?.toString() ?? '';

		if (!display_name || !username || !password || !confirm_password) {
			return fail(400, { display_name, username, error: 'All fields are required.' });
		}

		if (username.length < 3 || !/^[a-z0-9_]+$/.test(username)) {
			return fail(400, {
				display_name,
				username,
				error:
					'Username must be at least 3 characters and contain only letters, numbers, or underscores.'
			});
		}

		if (password.length < 8) {
			return fail(400, {
				display_name,
				username,
				error: 'Password must be at least 8 characters.'
			});
		}

		if (password !== confirm_password) {
			return fail(400, { display_name, username, error: 'Passwords do not match.' });
		}

		let response: Response;
		try {
			response = await api.auth.register({ display_name, username, password });
		} catch {
			return fail(503, {
				display_name,
				username,
				error: 'Could not reach the server. Try again later.'
			});
		}

		if (!response.ok) {
			let message = 'Registration failed.';
			try {
				const data = await response.json();
				message = data.message || data.error || message;
			} catch {
				/* ignore */
			}
			return fail(response.status, { display_name, username, error: message });
		}

		redirect(302, '/login?registered=1');
	}
};
