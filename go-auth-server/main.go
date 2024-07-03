package main

import (
	"gamebuddy/auth/db"
	"gamebuddy/auth/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	cors "github.com/rs/cors/wrapper/gin"
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

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if ctx.Request.Method == "OPTIONS" {
// 			ctx.AbortWithStatus(204)
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

func main() {

	r := gin.Default()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	})

	r.Use(c)

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
