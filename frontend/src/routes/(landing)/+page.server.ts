import { fail, redirect } from '@sveltejs/kit';
import type { RequestEvent } from './$types.js';
import { supabase } from '$lib/supabaseClient.js';

/** @type {import('./$types').Actions} */
export const actions = {
	register: async (event: RequestEvent) => {
		// TODO register the user
		const data = await event.request.formData();
		const name = data.get('name') as string;
		const email = data.get('email') as string;
		const message = data.get('message') as string;

		const { error } = await supabase.from('waitlist').insert({ name, email, message }).select();

		if (error) {
			return fail(500, { error: error.message });
		}

		return redirect(302, '/thank-you');
	}
};
