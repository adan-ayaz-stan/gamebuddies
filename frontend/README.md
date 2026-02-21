# GameBuddies — Frontend

SvelteKit 5 client for the GameBuddies platform.

## Tech stack

| Concern | Library |
|---|---|
| Framework | [SvelteKit 5](https://kit.svelte.dev) + TypeScript |
| Styling | Tailwind CSS v4 |
| UI primitives | [bits-ui](https://bits-ui.com) (Popover, Dialog, …) |
| Icons | [@lucide/svelte](https://lucide.dev) |
| Package manager | pnpm |
| Testing | Vitest (unit) + Playwright (e2e) |

## Project structure

```
frontend/
├── src/
│   ├── hooks.server.ts          # Auth guard: reads JWT cookie → locals.user
│   ├── lib/
│   │   ├── components/          # Shared UI (BorderBitsCard, dashboard panels, …)
│   │   ├── server/
│   │   │   ├── api.ts           # Server-side fetch helpers (all backend calls)
│   │   │   └── auth.ts          # Token refresh logic
│   │   ├── stores/
│   │   │   ├── auth.svelte.ts   # Reactive auth state
│   │   │   └── matchmaking.svelte.ts  # WebSocket matchmaking store
│   │   └── types/auth.ts        # TypeScript interfaces (User, Friend, Stats, …)
│   └── routes/
│       ├── +layout.svelte       # Root layout
│       ├── +page.svelte         # Landing page
│       ├── login/               # Sign-in page
│       ├── register/            # Sign-up page
│       ├── logout/              # Session teardown
│       ├── profile/             # Profile editor + avatar upload
│       └── dashboard/
│           ├── +layout.svelte   # Dashboard shell (sidebar, nav, friends list)
│           ├── +page.svelte     # Stats & overview
│           ├── matchmaking/     # Game selection + real-time queue
│           ├── friends/         # Send requests, accept/decline, friends list
│           └── messages/        # (placeholder)
├── static/
│   ├── fonts/                   # Big Shoulders + Iceland variable fonts
│   ├── sounds/                  # UI sound effects
│   └── svgs/
├── .env.example
├── package.json
└── svelte.config.js
```

## Quick start

### Prerequisites

- Node.js 20+
- pnpm — `npm i -g pnpm`
- Backend running at `http://localhost:8080`

### 1. Environment

```bash
cp .env.example .env.local
# VITE_API_URL=http://localhost:8080  (default, change if needed)
```

### 2. Install & run

```bash
pnpm install
pnpm dev
# http://localhost:5173
```

### 3. Build for production

```bash
pnpm build
pnpm preview
```

## Key pages

| Route | Description |
|---|---|
| `/` | Landing page |
| `/register` | Create account |
| `/login` | Sign in |
| `/dashboard` | Stats overview (platform time, matches, sessions) |
| `/dashboard/matchmaking` | Select games → join real-time matchmaking queue |
| `/dashboard/friends` | Send friend requests, accept/decline incoming, manage list |
| `/profile` | Edit display name, upload avatar |

## Auth architecture

- All tokens stay in **HttpOnly cookies** — never exposed to JavaScript
- `hooks.server.ts` validates the access-token cookie on every request and populates `locals.user`; transparently refreshes the token if expired
- Protected routes redirect to `/login` when `locals.user` is absent
- Server-side `load` functions and form `actions` handle all API calls — no client-side token management needed

## Matchmaking

The `matchmakingStore` (`lib/stores/matchmaking.svelte.ts`) opens a WebSocket to `ws://localhost:8080/v1/ws/matchmaking` and exposes reactive state (`status`, `match`, `queueSize`). The matchmaking page `connects` on mount and `disconnects` on destroy.

## Testing

```bash
pnpm test          # Vitest unit tests
pnpm test:e2e      # Playwright end-to-end tests
pnpm check         # svelte-check type validation
```
