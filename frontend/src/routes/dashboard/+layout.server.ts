import { api } from '$lib/server/api';
import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals, cookies }) => {
	// Auth guard – redirect to login if not authenticated
	if (!locals.user) {
		throw redirect(302, '/login');
	}

	const accessToken = cookies.get('gm_access_token') ?? '';

	// Load friends list and pending request count (non-fatal if it fails)
	const [friends, requests] = await Promise.all([
		api.users.getFriends(accessToken),
		api.users.getFriendRequests(accessToken)
	]);

	return {
		user: locals.user,
		friends,
		pendingRequestCount: requests.length
	};
};
