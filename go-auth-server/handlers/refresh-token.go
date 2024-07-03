package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gamebuddy/auth/db"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RefreshTokenRoute(ctx *gin.Context) {
	// when a user logs in, a cookie with the name of 'gamebuddy' is attached to the user's browser

	database := db.GetDB()

	// check for the 'gamebuddy' cookie
	authCookie, err := ctx.Cookie("gamebuddy")

	if err != nil || authCookie == "" {
		fmt.Println("no 'gamebuddy' cookie found")
		ctx.JSON(401, "Unauthorized, Missing Token.")
		return
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(authCookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err != nil {
		// remove the 'gamebuddy' cookie
		ctx.SetCookie("gamebuddy", "", -1, "/", "localhost", false, true)
		fmt.Println("invalid token")
		ctx.JSON(401, "Unauthorized, Invalid Token.")
		return
	}

	// get the token from the 'access_token' field in the claims
	refresh_token := claims["access_token"].(string)
	user_email := claims["email"].(string)

	// Get the user from the database
	var dbUser db.User
	database.Table("users").Where("email = ?", user_email).First(&dbUser)

	var dbRefreshToken db.RefreshTokens

	// check if user exists
	if dbUser.ID == "" {
		ctx.JSON(401, "Unauthorized, User not found.")
		fmt.Println("User not found")
		return
	}

	// check the token against the database record of 'refresh_tokens'
	database.Table("refresh_tokens").Find(&dbRefreshToken, "token = ?", refresh_token)

	if dbRefreshToken.Revoked {
		fmt.Println("Refresh token has been revoked")
		ctx.JSON(401, "Unauthorized, Refresh token has been revoked.")
		return
	}

	if dbRefreshToken.Token != refresh_token {
		fmt.Println("Invalid refresh token")
		ctx.JSON(401, "Unauthorized, Invalid refresh token.")
		return
	}

	// if it is valid, we generate a new token, which translates into a new auth cookie
	// Generate random hash string for refresh token
	hash := sha256.Sum256([]byte(fmt.Sprintf("User email: %s  Time: %d", user_email, time.Now().UnixNano())))
	access_token := hex.EncodeToString(hash[:])

	// Create authorization token through jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":        user_email,
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

	var session db.Sessions

	// find session for user and update it
	database.Table("sessions").Where("user_id = ?", dbUser.ID).First(&session)

	if session.ID == 0 {
		database.Table("sessions").Create(&db.Sessions{
			UserId:      dbUser.ID,
			RefreshedAt: &now,
			UserAgent:   &agent,
			Ip:          &ctx.Request.RemoteAddr,
		})
	} else {
		session.UserAgent = &agent
		session.RefreshedAt = &now
		database.Table("sessions").Save(&session)
	}

	database.Table("refresh_tokens").Create(&db.RefreshTokens{
		UserId:    dbUser.ID,
		Token:     access_token,
		SessionID: uint(session.ID),
		Parent:    &refresh_token,
		Revoked:   false,
	})

	// we send back the response which will be the user information (email, name) along with the new cookie
	// Assign token to cookie
	ctx.SetCookie("gamebuddy", signedString, 36000, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"email": dbUser.Email,
	})
}
