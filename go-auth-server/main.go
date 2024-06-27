package main

import (
	"gamebuddy/auth/handlers"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Temporarily putting structs here
type User struct {
	gorm.Model
	id                   string `gorm:"primaryKey"`
	instance_id          string
	aud                  string
	role                 string
	email                string
	encrypted_password   string
	email_confirmed_at   *time.Time
	confirmation_token   *string
	confirmation_sent_at *time.Time
	recovery_token       *string
	recovery_sent_at     *time.Time
	last_sign_in_at      *time.Time
	banned_until         *time.Time
}

// This is an authentication server. This server will connect to a database that will store the users.

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		api.GET("/heartbeat", handlers.HeartbeatRoute)
		api.POST("/login", handlers.LoginRoute)
		api.POST("/register", handlers.RegisterRoute)
		api.POST("/logout", handlers.LogoutRoute)
		api.POST("/refresh", handlers.RefreshRoute)
		api.POST("/verify-email", handlers.VerifyEmailRoute)
		api.POST("/reset-password", handlers.ResetPasswordRoute)
	}

	r.Run()
}
