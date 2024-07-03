package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gamebuddy/auth/db"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email             string `json:"email" validate:"required,email"`
	EncryptedPassword string `json:"encrypted_password" validate:"required"`
	RedirectUrl       string `json:"redirect_url" validate:"required"`
	OriginUrl         string `json:"origin_url" validate:"required"`
}

func LoginRoute(ctx *gin.Context) {

	// Check first if user is already logged in

	if _, err := ctx.Cookie("gamebuddy-auth"); err == nil {
		// redirect to refresh-token route - temporarily
		ctx.Redirect(302, "/api/v1/refresh")
		return
	}

	var request LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Input Validation Failed",
		})
		return
	}

	database := db.GetDB()

	var existingUser db.User
	database.Table("users").Where("email = ?", request.Email).First(&existingUser)

	if existingUser.ID == "" {
		ctx.JSON(401, gin.H{
			"message": "Invalid email.",
			"field":   "email",
		})
		return
	}

	// Hash password
	data := []byte(request.EncryptedPassword)
	hashedInputPass := sha256.Sum256(data)
	hashString := hex.EncodeToString(hashedInputPass[:])

	// Compare passwords

	if hashString != existingUser.EncryptedPassword {
		ctx.JSON(401, gin.H{
			"message": "Invalid password.",
			"field":   "password",
		})
		return
	}

	// Generate random hash string for refresh token
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	access_token := hex.EncodeToString(hash[:])

	// Generate JWT token
	// Create authorization token through jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":        request.Email,
		"access_token": access_token,
	})

	jwtSecret := os.Getenv("JWT_SECRET")

	signedString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		// Log error
		fmt.Println(fmt.Errorf("error while generating token: %v", err))

		ctx.JSON(500, gin.H{
			"message": "Internal server error. Error while generating token.",
		})
		return
	}

	now := time.Now()
	agent := ctx.Request.UserAgent()

	session := db.Sessions{
		UserId:      existingUser.ID,
		RefreshedAt: &now,
		UserAgent:   &agent,
		Ip:          &ctx.Request.RemoteAddr,
	}

	// Create a session
	database.Table("sessions").Create(&session)

	database.Table("refresh_tokens").Create(&db.RefreshTokens{
		UserId:    existingUser.ID,
		Token:     access_token,
		SessionID: uint(session.ID),
		Revoked:   false,
	})

	// Assign token to cookie
	ctx.SetCookie("gamebuddy", signedString, 36000, "/", "*", false, true)
	ctx.JSON(200, "OK")
}
