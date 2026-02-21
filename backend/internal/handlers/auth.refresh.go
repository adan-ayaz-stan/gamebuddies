package handlers

import (
	"net/http"
	"time"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RefreshTokenHandler(c *gin.Context) {
	refresh_token, err := c.Cookie("gm_refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token not found"})
		return
	}

	db := h.DB

	// O(1) lookup: parse user ID embedded in token format "<userID>:<secret>"
	userID, err := utils.ParseUserIDFromRefreshToken(refresh_token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	var user models.User
	if result := db.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	if user.RefreshTokenHash == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	// Verify the token against the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.RefreshTokenHash), []byte(refresh_token)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	// Check if the token is expired
	if time.Now().After(user.RefreshTokenExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token has expired"})
		return
	}

	// Token is valid - generate new access token
	jwtToken, accessExpiresAt, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token. Error: " + err.Error()})
		return
	}

	// Set new access token cookie
	accessMaxAge := int(time.Until(accessExpiresAt).Seconds())
	c.SetCookie(
		"gm_access_token",
		jwtToken,
		accessMaxAge,
		"/",
		"",
		false,
		true,
	)

	// Generate new refresh token (token rotation)
	plainRefreshToken, hashedRefreshToken, refreshExpiresAt, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token. Error: " + err.Error()})
		return
	}

	// Update user with new refresh token
	user.RefreshTokenHash = hashedRefreshToken
	user.RefreshTokenExpiresAt = refreshExpiresAt

	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new refresh token to database"})
		return
	}

	// Set new refresh token cookie
	refreshMaxAge := int(time.Until(refreshExpiresAt).Seconds())
	c.SetCookie(
		"gm_refresh_token",
		plainRefreshToken,
		refreshMaxAge,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tokens refreshed successfully",
	})
}
