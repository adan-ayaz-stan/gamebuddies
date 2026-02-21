package handlers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
)

// UpdateProfileHandler PATCH /v1/profile – update display name.
func (h *Handler) UpdateProfileHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	var req struct {
		DisplayName string `json:"display_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.DisplayName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "display_name is required"})
		return
	}

	db := h.DB
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		return
	}

	profile.DisplayName = strings.TrimSpace(req.DisplayName)
	if err := db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Profile updated",
		"display_name": profile.DisplayName,
		"avatar_url":   profile.AvatarURL,
	})
}

// UploadAvatarHandler POST /v1/profile/avatar – upload a profile picture to MinIO.
func (h *Handler) UploadAvatarHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	if h.Storage == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "File storage not configured"})
		return
	}

	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "avatar file is required"})
		return
	}
	defer file.Close()

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]string{".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".png": "image/png", ".webp": "image/webp", ".gif": "image/gif"}
	contentType, ok := allowed[ext]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file type. Allowed: jpg, png, webp, gif"})
		return
	}

	// Max 5 MB
	const maxSize = 5 << 20
	if header.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File too large. Maximum 5 MB"})
		return
	}

	objectName := fmt.Sprintf("avatars/user-%d%s", userID, ext)

	avatarURL, err := h.Storage.UploadAvatar(context.Background(), objectName, file, header.Size, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Upload failed: " + err.Error()})
		return
	}

	db := h.DB
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		return
	}

	profile.AvatarURL = avatarURL
	if err := db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save avatar URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Avatar uploaded successfully",
		"avatar_url": avatarURL,
	})
}
