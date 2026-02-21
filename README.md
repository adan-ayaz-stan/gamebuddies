# GameBuddies

A full-stack gaming social platform where players can find teammates, track stats, and manage their gaming crew.

## What it does

- **Auth** — JWT-based sign-up / sign-in with refresh tokens and server-side HttpOnly cookies
- **Matchmaking** — real-time WebSocket queue; join with your preferred games and get matched with other players
- **Friends** — mutual-consent friend system: send a request → other player accepts → friendship established. Includes incoming-request badges and online/offline presence in the sidebar
- **Profile** — display name, avatar upload (stored in MinIO), platform-time tracker, match/session stats
- **Games** — game catalogue seeded into Postgres, cover images hosted in MinIO

## Stack

| Layer | Technology |
|---|---|
| Frontend | SvelteKit 5, TypeScript, Tailwind CSS |
| Backend | Go (Gin), GORM, PostgreSQL |
| Realtime | WebSockets (Gorilla WS) |
| Storage | MinIO (S3-compatible) |
| Migrations | Goose |

## Project structure

```
gamebuddies/
├── backend/          # Go API server
│   ├── cmd/
│   │   ├── api/          # Server entry point
│   │   └── seed-games/   # CLI: uploads game images to MinIO
│   ├── internal/
│   │   ├── handlers/     # HTTP handlers & routes
│   │   ├── matchmaking/  # WebSocket hub & client
│   │   ├── models/       # GORM models
│   │   ├── storage/      # MinIO client
│   │   └── utils/        # JWT, validation helpers
│   ├── migrations/       # Goose SQL migrations
│   └── static/games/     # Game cover images
└── frontend/         # SvelteKit app
    └── src/
        ├── lib/          # Shared components, stores, API client
        └── routes/       # File-based pages (dashboard, auth, profile)
```

## Quick start

### Prerequisites

- Go 1.24+
- Node.js 20+ / pnpm
- PostgreSQL
- MinIO (or any S3-compatible store)
- [Goose](https://github.com/pressly/goose) CLI

### 1. Backend

```bash
cd backend
cp .env.example .env   # fill in your values
go run ./cmd/api/
```

**Run migrations:**
```bash
goose -dir migrations postgres "<your DSN>" up
```

**Seed game images into MinIO:**
```bash
go run ./cmd/seed-games/
```

### 2. Frontend

```bash
cd frontend
cp .env.example .env.local   # set VITE_API_URL if needed
pnpm install
pnpm dev
```

App runs at `http://localhost:5173`, API at `http://localhost:8080`.
