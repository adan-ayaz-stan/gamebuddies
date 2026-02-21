import { api } from '$lib/server/api';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	const accessToken = cookies.get('gm_access_token') ?? '';

	const [friends, requests] = await Promise.all([
		api.users.getFriends(accessToken),
		api.users.getFriendRequests(accessToken)
	]);

	return { friends, requests };
};

export const actions: Actions = {
	// Send a friend request
	sendRequest: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const username = form.get('username')?.toString().trim();
		if (!username) return fail(400, { sendError: 'Username is required' });

		const resp = await api.users.sendFriendRequest(accessToken, username);
		const body = await resp.json().catch(() => ({}));
		if (!resp.ok) {
			return fail(resp.status, { sendError: body.message ?? 'Could not send request' });
		}
		return { sendOk: true };
	},

	// Accept an incoming request
	accept: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const id = Number(form.get('id'));
		if (!id) return fail(400, { actionError: 'Invalid request' });

		const resp = await api.users.acceptFriendRequest(accessToken, id);
		if (!resp.ok) {
			const body = await resp.json().catch(() => ({}));
			return fail(resp.status, { actionError: body.message ?? 'Failed to accept' });
		}
		return { acceptOk: true };
	},

	// Decline an incoming request
	decline: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const id = Number(form.get('id'));
		if (!id) return fail(400, { actionError: 'Invalid request' });

		const resp = await api.users.declineFriendRequest(accessToken, id);
		if (!resp.ok) {
			const body = await resp.json().catch(() => ({}));
			return fail(resp.status, { actionError: body.message ?? 'Failed to decline' });
		}
		return { declineOk: true };
	},

	// Remove an existing friend
	remove: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const id = Number(form.get('id'));
		if (!id) return fail(400, { actionError: 'Invalid friend ID' });

		const resp = await api.users.removeFriend(accessToken, id);
		if (!resp.ok) {
			const body = await resp.json().catch(() => ({}));
			return fail(resp.status, { actionError: body.message ?? 'Failed to remove friend' });
		}
		return { removeOk: true };
	}
};
