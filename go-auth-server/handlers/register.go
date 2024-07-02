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

type RegisterRequest struct {
	Email             string `json:"email" validate:"required,email"`
	EncryptedPassword string `json:"encrypted_password" validate:"required"`
	RedirectUrl       string `json:"redirect_url" validate:"required"`
}

// I have not used this function because request is being used in the function below

// func ValidateRegisterRequest(ctx *gin.Context) bool {
// 	var request RegisterRequest
// 	if err := ctx.BindJSON(&request); err != nil {
// 		return false
// 	}
// 	validate := validator.New()
// 	err := validate.Struct(request)
// 	return err == nil
// }

// type MailRequest struct {
// 	Email            string `validate:"required,email"`
// 	VerificationCode string `validate:"required"`
// }

// func SendMail(recipient string, verificationCode string) *time.Time {

// 	// Set up auth
// 	auth := smtp.PlainAuth("", "spitfire", "password", "smtp.gmail.com")

// 	validate := validator.New()

// 	request := MailRequest{
// 		Email:            recipient,
// 		VerificationCode: verificationCode,
// 	}

// 	err := validate.Struct(request)

// 	if err != nil {
// 		fmt.Println("Mail request validation error: ", err)
// 		return nil
// 	}

// 	// We send the mail here

// 	now := time.Now()
// 	return &now
// }

func RegisterRoute(ctx *gin.Context) {

	database := db.GetDB()

	// API INPUT VALIDATION
	var request RegisterRequest

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

	// CHECK FOR EXISTING RECORD

	var existingUser db.User

	database.Table("users").Where("email = ?", request.Email).First(&existingUser)

	if existingUser.ID != "" {
		ctx.JSON(400, gin.H{
			"message": "Email already registered",
		})
		return
	}

	// Generate verification code

	// verificationCode := uuid.New().String()
	// now := time.Now()

	user := db.User{
		InstanceId:        "000-000-000",
		Aud:               "authenticated",
		Role:              "authenticated",
		Email:             request.Email,
		EncryptedPassword: request.EncryptedPassword,
		// ConfirmationToken:  &verificationCode,
		// ConfirmationSentAt: &now, // This is a zero time which is the same as time.Time{}.
	}

	// Register user in the database
	database.Table("users").Create(&user)

	// Send an email to the user with the verfication code
	// TURNED OFF BECAUSE NO CORPORATE EMAIL SERVER AVAILABLE

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

	fmt.Printf("User Object: %+v\n", user)

	session := db.Sessions{
		UserId:      user.ID,
		RefreshedAt: &now,
		UserAgent:   &agent,
		Ip:          &ctx.Request.RemoteAddr,
	}

	// Create a session
	database.Table("sessions").Create(&session)

	database.Table("refresh_tokens").Create(&db.RefreshTokens{
		UserId:    user.ID,
		Token:     access_token,
		SessionID: uint(session.ID),
		Revoked:   false,
	})

	// Assign token to cookie
	ctx.SetCookie("gamebuddy-auth", signedString, 36000, "/", "", false, true)
	ctx.JSON(200, "OK")
}
