package utils

import (
	"os"
	"time"

	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next() // Lewati pengecekan untuk OPTIONS request
			return
		}
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, APIResponse{Success: false, Message: "Authorization header is missing"})
			c.Abort()
			return
		}

		// Format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, APIResponse{Success: false, Message: "Token format invalid"})
			c.Abort()
			return
		}

		secretKey := os.Getenv("SECRET_KEY")

		// token verification
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		// if token invalid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, APIResponse{Success: false, Message: "Invalid token"})
			c.Abort()
			return
		}

		// token valid, set userID to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["userID"])
		}

		c.Next()
	}
}
