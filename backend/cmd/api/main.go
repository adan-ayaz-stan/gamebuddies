package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/db"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/handlers"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/matchmaking"
	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is live and healthy."})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure it exists in the project root.")
	}

	router := gin.Default()

	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	matchmakingHub := matchmaking.NewHub(database)
	go matchmakingHub.Run()

	// Initialise MinIO storage (best-effort – log warning if unavailable)
	store, err := storage.NewClient()
	if err != nil {
		log.Printf("WARNING: MinIO storage unavailable: %v", err)
		store = nil
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", HealthCheck)

	v1 := router.Group("/v1")
	h := handlers.NewHandler(database, matchmakingHub, store)

	h.SetupRoutes(v1)

	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
