package handlers

import (
	"net/http"
	"time"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) LogoutHandler(c *gin.Context) {
	// Support logout even with expired access token by also checking refresh token
	var userID uint

	accessToken, accErr := c.Cookie("gm_access_token")
	if accErr == nil {
		claims, err := utils.ValidateAccessToken(accessToken)
		if err == nil {
			userID = claims.UserID
		}
	}

	if userID == 0 {
		// Try to identify via refresh token (embedded user ID)
		refreshToken, rErr := c.Cookie("gm_refresh_token")
		if rErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}
		parsedID, pErr := utils.ParseUserIDFromRefreshToken(refreshToken)
		if pErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}
		userID = parsedID
	}

	db := h.DB

	// Find the user
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		// Still clear cookies even if user not found
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
		return
	}

	// Revoke refresh token in database
	user.RefreshTokenHash = ""
	user.RefreshTokenExpiresAt = time.Now().Add(-24 * time.Hour)

	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke refresh token"})
		return
	}

	// Close all active sessions for this user
	now := time.Now()
	db.Model(&models.UserSession{}).
		Where("user_id = ? AND ended_at IS NULL", user.ID).
		Update("ended_at", now)

	// Clear access token cookie
	c.SetCookie(
		"gm_access_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	// Clear refresh token cookie
	c.SetCookie(
		"gm_refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
