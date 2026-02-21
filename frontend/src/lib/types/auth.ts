/** Authenticated user – returned by GET /v1/profile and stored in locals.user */
export interface User {
	id: number;
	username: string;
	display_name: string;
	avatar_url?: string;
}

/** Shape of the user object nested inside the login response. */
export interface LoginUser {
	id: number;
	username: string;
	profile: {
		id: number;
		displayName: string;
		avatarURL?: string;
	};
}

/** Shape of POST /v1/auth/sign-in response body. */
export interface LoginResponse {
	message: string;
	user: LoginUser;
}

/** Shape of POST /v1/auth/sign-up response body. */
export interface RegisterResponse {
	message: string;
	user_id: number;
}

export interface LoginRequest {
	username: string;
	password: string;
}

export interface RegisterRequest {
	display_name: string;
	username: string;
	password: string;
}

/** Friend as returned by GET /v1/friends */
export interface Friend {
	id: number;
	username: string;
	display_name: string;
	avatar_url: string;
	is_online: boolean;
}

/** Incoming friend request as returned by GET /v1/friends/requests */
export interface FriendRequest {
	id: number;
	requester_id: number;
	username: string;
	display_name: string;
	avatar_url: string;
}

/** Stats as returned by GET /v1/stats */
export interface Stats {
	total_seconds: number;
	total_minutes: number;
	total_hours: number;
	total_sessions: number;
	total_matches: number;
}

export interface ApiError {
	error?: string;
	message?: string;
}
