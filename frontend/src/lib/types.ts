export type ColorVariant = 'red' | 'blue' | 'green' | 'yellow' | 'purple' | 'pink';

// API Types
export interface Profile {
	id: string;
	username: string;
	display_name: string;
	language: string;
	profile_picture_url?: string;
	mmr_rating: number;
	games_played: number;
	total_playtime: number;
	created_at: string;
	updated_at: string;
	is_online: boolean;
	last_seen: string;
}

export interface Game {
	id: string;
	name: string;
	category: string;
	min_players: number;
	max_players: number;
	is_active: boolean;
	created_at: string;
}

export interface RoomMember {
	id: string;
	room_id: string;
	user_id: string;
	username: string;
	display_name: string;
	profile: Profile;
	joined_at: string;
	left_at?: string;
	is_active: boolean;
	is_muted: boolean;
	in_voice: boolean;
	latency: number;
}

export interface Room {
	id: string;
	name: string;
	room_code: string;
	game_id: string;
	game: Game;
	game_mode: string;
	owner_id: string;
	status: 'waiting' | 'active' | 'dissolved';
	max_capacity: number;
	max_members: number;
	current_players: number;
	language: string;
	avg_latency: number;
	created_at: string;
	started_at?: string;
	ended_at?: string;
	members: RoomMember[];
}

export interface ChatMessage {
	id: string;
	room_id: string;
	user_id: string;
	sender_display_name: string;
	profile: {
		username: string;
		display_name: string;
		profile_picture_url?: string;
	};
	content: string;
	message: string;
	timestamp: string;
	sent_at: string;
	is_system_message: boolean;
}
