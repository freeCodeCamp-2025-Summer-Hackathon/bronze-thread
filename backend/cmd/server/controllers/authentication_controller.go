package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Type for the payload while signin.
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signin(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user db.User
	result := db.DB.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		// User not found, create new user
		user = db.User{
			Password: req.Password,
			Email:    req.Email,
		}
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// "Login" logic: For a real app, you'd set a cookie or JWT token here.
	if user.Password != req.Password {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password is not correct!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"name":  user.Name,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	jwt_secret_key := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwt_secret_key))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating auth token.",
			"error":   err,
		})
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed in successfully",
		"user":    user,
	})
}
