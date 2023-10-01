package middleware

import (
	"net/http"
	"v_com/auth"
	"v_com/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Extract the JWT token from the Authorization header
	tokenString := utils.ExtractTokenFromHeader(c.Request)

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Validate the JWT token
	claims, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Add the user ID to the request context
	c.Set("UserID", claims.UserID)
	c.Next()
}
