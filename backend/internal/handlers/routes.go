package handlers

import (
	"net/http"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/matchmaking"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB             *gorm.DB
	MatchmakingHub *matchmaking.Hub
	Storage        *storage.Client
}

func NewHandler(db *gorm.DB, hub *matchmaking.Hub, store *storage.Client) *Handler {
	return &Handler{DB: db, MatchmakingHub: hub, Storage: store}
}

func (h *Handler) indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, from the index route"})
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/", h.indexHandler)

	auth := router.Group("/auth")
	auth.POST("/sign-in", h.LoginHandler)
	auth.POST("/sign-up", h.RegisterHandler)
	auth.POST("/refresh", h.RefreshTokenHandler)
	auth.POST("/logout", h.LogoutHandler)

	protected := router.Group("")
	protected.Use(AuthMiddleware())

	// User
	protected.GET("/profile", h.GetUserProfileHandler)
	protected.PATCH("/profile", h.UpdateProfileHandler)
	protected.POST("/profile/avatar", h.UploadAvatarHandler)

	// Friends
	protected.GET("/friends", h.GetFriendsHandler)
	protected.POST("/friends", h.SendFriendRequestHandler)
	protected.DELETE("/friends/:friendId", h.RemoveFriendHandler)

	// Friend requests
	protected.GET("/friends/requests", h.GetFriendRequestsHandler)
	protected.POST("/friends/requests/:id/accept", h.AcceptFriendRequestHandler)
	protected.POST("/friends/requests/:id/decline", h.DeclineFriendRequestHandler)

	// Stats
	protected.GET("/stats", h.GetStatsHandler)

	// Matchmaking WS
	protected.GET("/ws/matchmaking", h.MatchmakingWsHandler)

	// Public
	matchmakingGroup := router.Group("/matchmaking")
	matchmakingGroup.GET("/games", h.GetMatchmakingGamesHandler)
}
