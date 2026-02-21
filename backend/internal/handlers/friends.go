package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FriendDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	IsOnline    bool   `json:"is_online"`
}

type FriendRequestDTO struct {
	ID          uint   `json:"id"`
	RequesterID uint   `json:"requester_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
}

// ── GET /v1/friends ──────────────────────────────────────────────────────────
// Returns accepted friends only.
func (h *Handler) GetFriendsHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	db := h.DB

	var friendships []models.Friendship
	if err := db.
		Preload("Requester.Profile").
		Preload("Addressee.Profile").
		Where("(requester_id = ? OR addressee_id = ?) AND status = ?", userID, userID, models.FriendshipAccepted).
		Find(&friendships).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	onlineSet := h.MatchmakingHub.GetOnlineUserIDs()

	friends := make([]FriendDTO, 0, len(friendships))
	for _, f := range friendships {
		var friendUser models.User
		if f.RequesterID == userID {
			friendUser = f.Addressee
		} else {
			friendUser = f.Requester
		}

		_, isOnline := onlineSet[friendUser.ID]

		dto := FriendDTO{
			ID:       friendUser.ID,
			Username: friendUser.Username,
			IsOnline: isOnline,
		}
		if friendUser.Profile.ID != 0 {
			dto.DisplayName = friendUser.Profile.DisplayName
			dto.AvatarURL = friendUser.Profile.AvatarURL
		}
		friends = append(friends, dto)
	}

	c.JSON(http.StatusOK, friends)
}

// ── POST /v1/friends  { "username": "..." } ──────────────────────────────────
// Sends a friend request (creates a pending friendship).
func (h *Handler) SendFriendRequestHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	db := h.DB

	var req struct {
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username is required"})
		return
	}
	req.Username = strings.ToLower(strings.TrimSpace(req.Username))

	// Find target user
	var target models.User
	if err := db.Where("username = ?", req.Username).First(&target).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	if target.ID == userID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot send a friend request to yourself"})
		return
	}

	// Check if any friendship/request already exists (in either direction)
	var existing models.Friendship
	err := db.Where(
		"(requester_id = ? AND addressee_id = ?) OR (requester_id = ? AND addressee_id = ?)",
		userID, target.ID, target.ID, userID,
	).First(&existing).Error

	if err == nil {
		if existing.Status == models.FriendshipAccepted {
			c.JSON(http.StatusConflict, gin.H{"message": "Already friends"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"message": "Friend request already pending"})
		}
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	friendship := models.Friendship{
		RequesterID: userID,
		AddresseeID: target.ID,
		Status:      models.FriendshipPending,
	}
	if err := db.Create(&friendship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send friend request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Friend request sent", "request_id": friendship.ID})
}

// ── GET /v1/friends/requests ─────────────────────────────────────────────────
// Returns incoming pending friend requests for the authenticated user.
func (h *Handler) GetFriendRequestsHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	db := h.DB

	var friendships []models.Friendship
	if err := db.
		Preload("Requester.Profile").
		Where("addressee_id = ? AND status = ?", userID, models.FriendshipPending).
		Find(&friendships).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	requests := make([]FriendRequestDTO, 0, len(friendships))
	for _, f := range friendships {
		dto := FriendRequestDTO{
			ID:          f.ID,
			RequesterID: f.RequesterID,
			Username:    f.Requester.Username,
		}
		if f.Requester.Profile.ID != 0 {
			dto.DisplayName = f.Requester.Profile.DisplayName
			dto.AvatarURL = f.Requester.Profile.AvatarURL
		}
		requests = append(requests, dto)
	}

	c.JSON(http.StatusOK, requests)
}

// ── POST /v1/friends/requests/:id/accept ─────────────────────────────────────
// Accepts an incoming friend request.
func (h *Handler) AcceptFriendRequestHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	idStr := c.Param("id")
	requestID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request ID"})
		return
	}

	db := h.DB

	var friendship models.Friendship
	if err := db.First(&friendship, uint(requestID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Friend request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	// Only the addressee may accept
	if friendship.AddresseeID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not authorised to accept this request"})
		return
	}
	if friendship.Status != models.FriendshipPending {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is not pending"})
		return
	}

	if err := db.Model(&friendship).Update("status", models.FriendshipAccepted).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to accept request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request accepted"})
}

// ── POST /v1/friends/requests/:id/decline ────────────────────────────────────
// Declines (deletes) an incoming friend request.
func (h *Handler) DeclineFriendRequestHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	idStr := c.Param("id")
	requestID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request ID"})
		return
	}

	db := h.DB

	var friendship models.Friendship
	if err := db.First(&friendship, uint(requestID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Friend request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	// Only the addressee may decline; requester may also cancel their own outgoing request
	if friendship.AddresseeID != userID && friendship.RequesterID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not authorised"})
		return
	}
	if friendship.Status != models.FriendshipPending {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is not pending"})
		return
	}

	if err := db.Delete(&friendship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decline request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request declined"})
}

// ── DELETE /v1/friends/:friendId ─────────────────────────────────────────────
// Removes an accepted friendship.
func (h *Handler) RemoveFriendHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	friendIDStr := c.Param("friendId")
	friendID, err := strconv.ParseUint(friendIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid friend ID"})
		return
	}

	db := h.DB
	result := db.Where(
		"(requester_id = ? AND addressee_id = ?) OR (requester_id = ? AND addressee_id = ?)",
		userID, uint(friendID), uint(friendID), userID,
	).Delete(&models.Friendship{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Friendship not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend removed"})
}
