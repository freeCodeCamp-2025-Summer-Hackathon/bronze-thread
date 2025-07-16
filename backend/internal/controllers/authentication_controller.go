package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/database"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	// Check if a user with the given email already exists
	var existingUser models.User
	if database.DB.Where("email = ?", req.Email).First(&existingUser).Error == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "User with this email already exists"})
		return
	}

	// Hash the password from the request
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to hash password"})
		return
	}

	// Create a new user model with the hashed password
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	// Save the new user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with a success message and the new UserID
	c.JSON(http.StatusCreated, dto.RegisterResponse{Message: "User registered successfully", UserID: user.ID})

}

func Signin(c *gin.Context) {
	var req dto.LoginRequest

	// Bind the incoming JSON to the LoginRequest DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	// Find the user by email
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "User not found or invalid credentials"})
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	// Generate a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create token"})
		return
	}

	// Set the token in a secure, HTTP-only cookie
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login successful",
		Token:   tokenString,
		User: dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Bio:          user.Bio,
			AvatarURL:    user.AvatarURL,
			City:         user.City,
			State:        user.State,
			RatingAvg:    user.RatingAvg,
			TotalRatings: user.TotalRatings,
			TotalSwaps:   user.TotalSwaps,
			CreatedAt:    user.CreatedAt,
		},
	})
}
