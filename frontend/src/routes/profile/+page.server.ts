import { api } from '$lib/server/api';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (!locals.user) {
		throw redirect(302, '/login');
	}
	return { user: locals.user };
};

export const actions: Actions = {
	updateDisplayName: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const display_name = form.get('display_name')?.toString().trim();

		if (!display_name || display_name.length < 2) {
			return fail(400, { updateError: 'Display name must be at least 2 characters' });
		}
		if (display_name.length > 50) {
			return fail(400, { updateError: 'Display name must be 50 characters or less' });
		}

		const resp = await api.users.updateProfile(accessToken, { display_name });
		if (!resp.ok) {
			const body = await resp.json().catch(() => ({}));
			return fail(400, { updateError: (body as Record<string, string>).error ?? 'Update failed' });
		}
		return { updateOk: true };
	},

	uploadAvatar: async ({ cookies, request }) => {
		const accessToken = cookies.get('gm_access_token') ?? '';
		const form = await request.formData();
		const file = form.get('avatar');

		if (!file || !(file instanceof File)) {
			return fail(400, { avatarError: 'No file provided' });
		}
		if (file.size > 5 * 1024 * 1024) {
			return fail(400, { avatarError: 'File must be under 5 MB' });
		}
		const allowed = ['image/jpeg', 'image/png', 'image/webp', 'image/gif'];
		if (!allowed.includes(file.type)) {
			return fail(400, { avatarError: 'Only JPG, PNG, WEBP or GIF allowed' });
		}

		// Forward multipart to backend
		const forward = new FormData();
		forward.append('avatar', file, file.name);

		const resp = await fetch('http://localhost:8080/v1/profile/avatar', {
			method: 'POST',
			headers: { Cookie: `gm_access_token=${accessToken}` },
			body: forward
		});

		if (!resp.ok) {
			const body = await resp.json().catch(() => ({}));
			return fail(400, { avatarError: (body as Record<string, string>).error ?? 'Upload failed' });
		}
		return { avatarOk: true };
	}
};
