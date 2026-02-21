import { getRequestEvent } from '$app/server';
import type { User } from '$lib/types/auth';
import { redirect } from '@sveltejs/kit';

/**
 * Authentication guard - ensures user is logged in
 * Redirects to login page if not authenticated
 */
export function requireAuth(): User {
	const { locals, url } = getRequestEvent();

	if (!locals.user) {
		const redirectTo = url.pathname + url.search;
		redirect(307, `/login?redirectTo=${encodeURIComponent(redirectTo)}`);
	}

	return locals.user;
}

/**
 * Get current user from locals
 * Returns null if not authenticated
 */
export function getUser(): User | null {
	const { locals } = getRequestEvent();
	return locals.user ?? null;
}

/**
 * Role-based authorization guard
 * Note: Backend doesn't currently support roles
 * This is a placeholder for future implementation
 */
export function requireRole(_allowedRoles: string[]): User {
	const user = requireAuth();

	// Future implementation:
	// if (user.role && !allowedRoles.includes(user.role)) {
	// 	error(403, 'Insufficient permissions');
	// }

	return user;
}
