package handlers

import (
	"net/http"
	"time"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
)

// GetStatsHandler GET /v1/stats
func (h *Handler) GetStatsHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	db := h.DB

	// Fetch profile for match count
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Profile not found"})
		return
	}

	// Fetch all sessions
	var sessions []models.UserSession
	if err := db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	var totalSeconds int64
	for _, s := range sessions {
		end := time.Now()
		if s.EndedAt != nil {
			end = *s.EndedAt
		}
		dur := end.Sub(s.StartedAt)
		if dur > 0 {
			totalSeconds += int64(dur.Seconds())
		}
	}

	totalMinutes := totalSeconds / 60
	totalHours := totalMinutes / 60

	c.JSON(http.StatusOK, gin.H{
		"total_seconds":  totalSeconds,
		"total_minutes":  totalMinutes,
		"total_hours":    totalHours,
		"total_sessions": len(sessions),
		"total_matches":  profile.TotalMatches,
	})
}
