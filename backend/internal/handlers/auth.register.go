package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	var existingUser models.User

	db := h.DB

	req.Username = strings.ToLower(req.Username)

	result := db.Where("username = ?", req.Username).First(&existingUser)
	
	if result.Error != nil {

		// check if error is other than record not found
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error while checking username"})
			return
		}

	} else {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists."})
		return
	}

	// if we reach here, proceed with registration
	var user models.User

	if len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long"})
		return
	}

	err := user.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong while hashing the password"})
		return
	}

	user = models.User{
		Username: req.Username,
		Password: user.Password, // this password is attached when the user.HashPassword is called
	}

	result = db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	var userProfile models.Profile = models.Profile{
		DisplayName: req.DisplayName,
		User: &user,
	}

	result = db.Create(&userProfile)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}