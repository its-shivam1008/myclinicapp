package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"myclinic/internal/models"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // doctor or receptionist
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := DB.Where("username = ?", req.Username).First(&user).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		// Registering new user
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user = models.User{
			Username: req.Username,
			Password: string(hashedPassword),
			Role:     req.Role,
		}
		if err := DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	} else {
		// Loggin existing user
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}
	}

	// Generating JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}