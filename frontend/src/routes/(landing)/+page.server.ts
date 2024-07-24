import { redirect } from '@sveltejs/kit';
import type { RequestEvent } from './$types.js';

/** @type {import('./$types').Actions} */
export const actions = {
	waitlist_register: async (event: RequestEvent) => {
		// TODO register the user
		const data = await event.request.formData();
		const name = data.get('name') as string;
		const email = data.get('email') as string;

		console.log(name, email);

		redirect(303, '/thank-you');
	}
};
