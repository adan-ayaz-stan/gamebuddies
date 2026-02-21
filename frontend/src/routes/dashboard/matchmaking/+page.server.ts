import type { PageServerLoad } from './$types';

interface Game {
	id: number;
	title: string;
	subtitle: string;
	img_url: string;
}

export const load: PageServerLoad = async () => {
	try {
		const resp = await fetch('http://localhost:8080/v1/matchmaking/games');
		if (resp.ok) {
			const games = (await resp.json()) as Game[];
			return { games };
		}
	} catch {
		// Non-fatal – just show empty list
	}
	return { games: [] as Game[] };
};
