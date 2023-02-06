package middleware

import (
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/responses"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			responses.Error(http.StatusUnauthorized, "Missing token", c)
			c.Abort()
			return
		}

		// Cek apakah token ada di database dan masih valid
		userID, err := ValidateToken(tokenString)
		if err != nil {
			responses.Error(http.StatusUnauthorized, "Invalid token", c)
			c.Abort()
			return
		}

		// Set the user's ID to the user's ID in the token
		c.Set("userID", userID)
		c.Next()
	}
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.JwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	userID, ok := token.Claims.(jwt.MapClaims)["userID"]
	if !ok {
		return "", fmt.Errorf("Claim 'userID' not found in token")
	}

	return userID.(string), nil
}
