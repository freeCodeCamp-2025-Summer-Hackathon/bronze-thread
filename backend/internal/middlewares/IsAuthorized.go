package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized", Message: "Authentication token is missing"})
			c.Abort()
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized", Message: "Invalid authentication token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract the user ID ('sub' claim) from the token
			userIDFloat, ok := claims["sub"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized", Message: "Invalid token claims"})
				c.Abort()
				return
			}

			// Set just the user ID in the context for the next handlers
			c.Set("sub", uint(userIDFloat))

		} else {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized", Message: "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
