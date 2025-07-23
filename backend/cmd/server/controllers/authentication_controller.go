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

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user db.User
	result := db.DB.Where("email = ?", req.Email).First(&user)

	if result.Error == nil {
		// User exists
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists!"})
		return
	}

	user = db.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registration successful.",
		"user":    user,
	})
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
		// User not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found, register a new account!"})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password is not correct!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Username,
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

func SignOut(c *gin.Context) {
	_, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not signed in.",
		})
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed out successfully.",
	})
}

func GetSignedInUser(c *gin.Context) {
	userInterface, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not signed in.",
		})
		return
	}

	user, ok := userInterface.(db.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "ok",
		"data":    user,
	})
}
