package main

import (
	"gamebuddy/auth/db"
	"gamebuddy/auth/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var limiter = ratelimit.NewBucket(time.Second, 10)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func main() {

	r := gin.Default()
	db.InitDB()

	r.Use(RateLimiterMiddleware())

	api := r.Group("/api/v1")

	{
		api.GET("/heartbeat", handlers.HeartbeatRoute)
		api.POST("/login", handlers.LoginRoute)
		api.POST("/register", handlers.RegisterRoute)
		api.POST("/logout", handlers.LogoutRoute)
		api.GET("/refresh", handlers.RefreshTokenRoute)
		api.POST("/verify-email", handlers.VerifyEmailRoute)
		api.POST("/reset-password", handlers.ResetPasswordRoute)
	}

	r.Run()
}
