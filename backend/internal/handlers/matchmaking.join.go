package handlers

import (
	"net/http"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/matchmaking" // Use your new package path
	"github.com/gin-gonic/gin"
)


func (h *Handler) MatchmakingWsHandler(c *gin.Context) {
	// 1. Get UserID from your auth middleware
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal authentication error"})
		return
	}

	// 2. Just hand off to ServeWs.
	// NO checks here. The Hub handles everything.
	matchmaking.ServeWs(h.MatchmakingHub, c, userID)
}