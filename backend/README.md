# GameBuddies — Backend

Go API server powering the GameBuddies platform. Handles authentication, real-time matchmaking, the friends system, profile management, and game-image storage.

## Tech stack

| Concern | Library |
|---|---|
| HTTP framework | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io) + `gorm.io/driver/postgres` |
| Auth | `golang-jwt/jwt/v5` + `golang.org/x/crypto` |
| Real-time | [Gorilla WebSocket](https://github.com/gorilla/websocket) |
| Object storage | [minio-go v7](https://github.com/minio/minio-go) |
| Migrations | [Goose](https://github.com/pressly/goose) |
| Config | [godotenv](https://github.com/joho/godotenv) |

## Project structure

```
backend/
├── cmd/
│   ├── api/            # Server entry point
│   └── seed-games/     # CLI: uploads game covers to MinIO
├── internal/
│   ├── db/             # GORM database initialisation
│   ├── handlers/       # HTTP handlers + route registration
│   ├── matchmaking/    # WebSocket hub, client, message types
│   ├── models/         # GORM models (User, Profile, Friendship, Game, …)
│   ├── storage/        # MinIO client (avatars + game images)
│   └── utils/          # JWT helpers, refresh-token logic, validators
├── migrations/         # Goose SQL migrations (chronological)
├── static/
│   └── games/          # Game cover images served into MinIO
├── .env.example
├── go.mod
└── go.sum
```

## Quick start

### Prerequisites

- Go 1.24+
- PostgreSQL (default: `localhost:5432`, db `gm_db`, user `admin`)
- MinIO (default: `localhost:9000`)
- [Goose](https://github.com/pressly/goose) CLI — `go install github.com/pressly/goose/v3/cmd/goose@latest`

### 1. Environment

```bash
cp .env.example .env
# Edit .env — set JWT_SECRET, DB credentials, MinIO credentials
```

### 2. Run migrations

```bash
goose -dir migrations postgres "host=localhost user=admin password=admin123 dbname=gm_db port=5432 sslmode=disable" up
```

### 3. Seed game images into MinIO

```bash
go run ./cmd/seed-games/
# Uploads backend/static/games/*.jpg to the gm-public MinIO bucket
# Re-run after every DB reset
```

### 4. Start the server

```bash
go run ./cmd/api/
# Listening on http://localhost:8080
```

To reset the database and start fresh:

```bash
goose -dir migrations postgres "<DSN>" reset
goose -dir migrations postgres "<DSN>" up
go run ./cmd/seed-games/
```

## Environment variables

| Variable | Description | Default |
|---|---|---|
| `JWT_SECRET` | HS256 signing key — **change in production** | — |
| `JWT_EXPIRATION_HOURS` | Access token lifetime (hours) | `24` |
| `REFRESH_TOKEN_EXPIRY_HOURS` | Refresh token lifetime (hours) | `168` |
| `DATABASE_URL` | PostgreSQL DSN | see `.env.example` |
| `SERVER_PORT` | Gin listen address | `:8080` |
| `GOOSE_DRIVER` | Goose driver | `postgresql` |
| `GOOSE_DBSTRING` | Goose DSN | see `.env.example` |
| `GOOSE_MIGRATION_DIR` | Migration directory | `./migrations` |
| `MINIO_ENDPOINT` | MinIO host:port | `localhost:9000` |
| `MINIO_ACCESS_KEY` | MinIO access key | — |
| `MINIO_SECRET_KEY` | MinIO secret key | — |
| `MINIO_USE_SSL` | Use HTTPS for MinIO | `false` |

## API reference

All routes are prefixed `/v1`.

### Auth (public)

| Method | Path | Description |
|---|---|---|
| `POST` | `/v1/auth/sign-up` | Register — `{ display_name, username, password }` |
| `POST` | `/v1/auth/sign-in` | Login — `{ username, password }` — sets HttpOnly cookies |
| `POST` | `/v1/auth/refresh` | Refresh access token (reads `gm_refresh_token` cookie) |
| `POST` | `/v1/auth/logout` | Revoke refresh token |

### Profile (protected)

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/profile` | Get current user's profile |
| `PATCH` | `/v1/profile` | Update display name |
| `POST` | `/v1/profile/avatar` | Upload avatar (multipart) → stored in MinIO |

### Friends (protected)

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/friends` | List accepted friends with online status |
| `POST` | `/v1/friends` | Send a friend request — `{ username }` |
| `DELETE` | `/v1/friends/:friendId` | Remove an accepted friend |
| `GET` | `/v1/friends/requests` | List incoming pending requests |
| `POST` | `/v1/friends/requests/:id/accept` | Accept a pending request |
| `POST` | `/v1/friends/requests/:id/decline` | Decline / cancel a request |

### Stats (protected)

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/stats` | Session count, total time, match count |

### Matchmaking (protected)

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/ws/matchmaking` | WebSocket upgrade — join/leave queue |

### Public

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/matchmaking/games` | List available games (title, subtitle, img_url) |

## Authentication flow

1. `POST /v1/auth/sign-up` — create account
2. `POST /v1/auth/sign-in` — backend sets `gm_access_token` + `gm_refresh_token` HttpOnly cookies
3. Include the access-token cookie on every protected request
4. `POST /v1/auth/refresh` when the access token expires (24 h)
5. `POST /v1/auth/logout` to revoke the session

## MinIO buckets

| Bucket | Purpose | Access |
|---|---|---|
| `gamebuddies-avatars` | User avatar images | Private (pre-signed URLs) |
| `gm-public` | Game cover images | Public read |

## Database migrations

Migrations live in `migrations/` and are managed with Goose.

```bash
goose -dir migrations postgres "<DSN>" status   # check state
goose -dir migrations postgres "<DSN>" up       # apply all
goose -dir migrations postgres "<DSN>" reset    # roll back all
```

## Build

```bash
go build ./cmd/api/           # build server binary
go build ./cmd/seed-games/    # build seeder binary
go build ./...                # verify all packages compile
```


### 1. Clone and Setup

```bash
cd f:\Coding\gamebuddies\backend
```

### 2. Configure Environment

```bash
# Copy the example environment file
cp .env.example .env

# Edit .env and set your values (IMPORTANT: Change JWT_SECRET!)
```

### 3. Start PostgreSQL

Ensure your PostgreSQL database is running and accessible with the credentials in your `.env` file.

### 4. Run the Server

```bash
go run cmd/api/main.go
```

The server will start at `http://localhost:8080`

### 5. Test the API

**Option 1: Using PowerShell Script (Windows)**

```powershell
.\test_auth_api.ps1
```

**Option 2: Using Bash Script (Linux/Mac/WSL)**

```bash
chmod +x test_auth_api.sh
./test_auth_api.sh
```

**Option 3: Import Postman Collection**

- Import `postman_collection.json` into Postman
- Set the `base_url` variable to `http://localhost:8080/v1`
- Run the requests in order

**Option 4: Manual cURL**

```bash
# Register
curl -X POST http://localhost:8080/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass123"}'

# Login
curl -X POST http://localhost:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass123"}'

# Get profile (replace TOKEN with your access token)
curl -X GET http://localhost:8080/v1/auth/profile \
  -H "Authorization: Bearer TOKEN"
```

## 📚 Documentation

- **[AUTH_API.md](./AUTH_API.md)** - Complete API documentation with all endpoints
- **[IMPLEMENTATION_SUMMARY.md](./IMPLEMENTATION_SUMMARY.md)** - Technical implementation details
- **[.env.example](./.env.example)** - Environment variables configuration

## 🔑 Authentication Flow

1. **Register** → Get access token + refresh token
2. **Use access token** in `Authorization: Bearer <token>` header
3. **Refresh** when access token expires (24h default)
4. **Logout** to invalidate refresh token

## 📋 Available Endpoints

### Public (No Authentication)

- `POST /v1/auth/register` - Register new user
- `POST /v1/auth/login` - Login
- `POST /v1/auth/refresh` - Refresh token

### Protected (Requires Bearer Token)

- `POST /v1/auth/logout` - Logout
- `GET /v1/auth/me` - Get current user
- `GET /v1/auth/profile` - Get profile
- `PATCH /v1/auth/profile` - Update profile
- `PATCH /v1/auth/password` - Change password

## 🏗️ Project Structure

```
backend/
├── cmd/api/main.go              # Entry point
├── internal/
│   ├── config/                  # Configuration
│   ├── db/                      # Database
│   ├── handlers/                # HTTP handlers
│   ├── middleware/              # Auth middleware
│   ├── models/                  # Data models & DTOs
│   └── utils/                   # JWT utilities
├── .env.example                 # Environment template
├── AUTH_API.md                  # API documentation
├── IMPLEMENTATION_SUMMARY.md    # Implementation details
├── postman_collection.json      # Postman collection
├── test_auth_api.ps1           # PowerShell test script
└── test_auth_api.sh            # Bash test script
```

## 🔒 Security Notes

⚠️ **IMPORTANT for Production:**

1. Change `JWT_SECRET` to a strong random value
2. Use HTTPS only
3. Set appropriate token expiration times
4. Implement rate limiting
5. Add CORS configuration
6. Use environment-specific configurations
7. Enable database SSL/TLS

## 🛠️ Development

### Build

```bash
go build -o gamebuddies-api cmd/api/main.go
```

### Run

```bash
./gamebuddies-api
```

### Test Compilation

```bash
go build ./cmd/api
```

## 🐛 Troubleshooting

**Database connection failed:**

- Check PostgreSQL is running
- Verify DATABASE_URL in .env
- Ensure database exists and user has permissions

**Token validation failed:**

- Check JWT_SECRET is same as when token was created
- Verify token hasn't expired
- Ensure Authorization header format: `Bearer <token>`

**Build errors:**

- Run `go mod tidy` to clean dependencies
- Ensure Go 1.24+ is installed

## 📦 Dependencies

All dependencies are managed via `go.mod`:

- `gin-gonic/gin` - Web framework
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - PostgreSQL driver
- `golang-jwt/jwt/v5` - JWT implementation
- `golang.org/x/crypto` - Password hashing

## 🎯 Next Steps

1. ✅ Basic authentication is complete
2. 🔜 Add email verification
3. 🔜 Implement password reset
4. 🔜 Add OAuth2 providers
5. 🔜 Implement 2FA
6. 🔜 Add rate limiting
7. 🔜 Setup logging middleware
8. 🔜 Add API documentation (Swagger)

## 📄 License

MIT License

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

**Status:** ✅ Fully functional JWT authentication system
**Last Updated:** October 23, 2025
