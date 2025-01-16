package middlewares

import (
	"learn-go-gin/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next() // Lewati pengecekan untuk OPTIONS request
			return
		}
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.APIResponse{Success: false, Message: "Authorization header is missing"})
			c.Abort()
			return
		}

		// Format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, utils.APIResponse{Success: false, Message: "Token format invalid"})
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
			c.JSON(http.StatusUnauthorized, utils.APIResponse{Success: false, Message: "Invalid token"})
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
