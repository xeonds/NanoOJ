package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(requiredPermissionLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// check if the user has the required permission level
		userPermissionLevel := int(token.Claims.(jwt.MapClaims)["permission_level"].(float64))
		if userPermissionLevel < requiredPermissionLevel {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		// set user ID and permission level in request context
		userID := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))
		c.Set("user_id", userID)
		c.Set("permission_level", userPermissionLevel)

		c.Next()
	}
}
