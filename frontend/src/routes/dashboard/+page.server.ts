import { api } from '$lib/server/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	const accessToken = cookies.get('gm_access_token') ?? '';
	const stats = await api.users.getStats(accessToken);

	return { stats };
};
