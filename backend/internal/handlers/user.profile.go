package handlers

import (
	"errors"
	"net/http"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetUserProfileHandler(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse User ID in context"})
		return
	}

	db := h.DB

	var user models.User
	results := db.Preload("Profile").Where("id = ?", userID).First(&user)
	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error})
		return
	}

	c.JSON(200, gin.H{
		"user_id":      user.ID,
		"username":     user.Username,
		"display_name": user.Profile.DisplayName,
		"avatar_url":   user.Profile.AvatarURL,
	})
}
