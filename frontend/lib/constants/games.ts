export type Game = {
  id: string;
  title: string;
  banner: string;
  company: string;
};

export const games: Game[] = [
  {
    id: "game-uuid-001",
    title: "League of Legends",
    banner: "/images/league-of-legends-banner.jpg",
    company: "RIOT",
  },
  {
    id: "game-uuid-002",
    title: "Apex Legends",
    banner: "/images/apex-legends-banner.jpg",
    company: "????",
  },
  {
    id: "game-uuid-003",
    title: "Call Of Duty Warzone",
    banner: "/images/call-of-duty-warzone-banner.jpg",
    company: "TENFO",
  },
  {
    id: "game-uuid-004",
    title: "Valorant",
    banner: "/images/valorant-banner.jpg",
    company: "Riot Games",
  },
];
