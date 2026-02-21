package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	db := h.DB

	var user models.User
	result := db.Preload("Profile").Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error while checking username"})
		return
	}

	var checkPass bool = user.CheckPassword(req.Password)
	if !checkPass {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	// Generate access token
	jwtToken, expiresAt, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token. Error: " + err.Error()})
		return
	}

	maxAge := int(time.Until(expiresAt).Seconds())

	c.SetCookie("gm_access_token", jwtToken, maxAge, "/", "", false, true)

	plainRefreshToken, hashedRefreshToken, refreshExpiresAt, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token. Error: " + err.Error()})
		return
	}

	user.RefreshTokenHash = hashedRefreshToken
	user.RefreshTokenExpiresAt = refreshExpiresAt

	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save refresh token to database"})
		return
	}

	refreshMaxAge := int(time.Until(refreshExpiresAt).Seconds())
	c.SetCookie("gm_refresh_token", plainRefreshToken, refreshMaxAge, "/", "", false, true)

	// Track platform session
	session := models.UserSession{
		UserID:    user.ID,
		StartedAt: time.Now(),
	}
	db.Create(&session) // best-effort, don't fail login if this fails

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"profile": gin.H{
				"id":          user.Profile.ID,
				"displayName": user.Profile.DisplayName,
				"avatarURL":   user.Profile.AvatarURL,
			},
		},
	})
}
