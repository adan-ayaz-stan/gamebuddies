# GameBuddies Backend - Quick Start Guide

## 🚀 Quick Start

### Prerequisites

- Go 1.24 or higher
- PostgreSQL database
- Git

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
