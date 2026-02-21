package handlers

import (
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type GameResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImgURL   string `json:"img_url"`
}

func (h *Handler) GetMatchmakingGamesHandler(c *gin.Context) {
	var games []GameResponse

	db := h.DB

	results := db.Model(&models.Game{}).Select("id", "title", "subtitle", "img_url").Scan(&games)
	if results.Error != nil {
		c.JSON(500, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(200, games)
}
