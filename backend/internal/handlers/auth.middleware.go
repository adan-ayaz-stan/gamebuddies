package handlers

import (
	"net/http"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("gm_access_token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired access token. Error: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
