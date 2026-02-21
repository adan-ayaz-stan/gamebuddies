export const load = async ({ locals }: { locals: App.Locals }) => {
	// Make user data available to all pages
	return {
		user: locals.user
	};
};
