type RoomStatus = "OPEN" | "CLOSED" | "FULL" | "PRIVATE";

type Room = {
  name: string;
  karma_points: number;
  capacity: number;
  status: RoomStatus;
};

type Game = {
  name: string;
  banner: string;
};

export type TRoom = {
  room: Room;
  game: Game;
};
