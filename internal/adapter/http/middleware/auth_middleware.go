package middleware

import (
	"net/http"
	"strings"

	"kochappi/internal/application/port"

	"github.com/gin-gonic/gin"
)

const (
	CONTEXT_KEY_USER_ID = "user_id"
	CONTEXT_KEY_ROLE    = "role"
)

func AuthMiddleware(tokenProvider port.TokenProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			return
		}

		userID, role, err := tokenProvider.ValidateAccessToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(CONTEXT_KEY_USER_ID, userID)
		c.Set(CONTEXT_KEY_ROLE, role)
		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) int {
	userID, _ := c.Get(CONTEXT_KEY_USER_ID)
	return userID.(int)
}

func GetRoleFromContext(c *gin.Context) string {
	role, _ := c.Get(CONTEXT_KEY_ROLE)
	return role.(string)
}

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := GetRoleFromContext(c)
		for _, r := range allowedRoles {
			if role == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions", "code": "FORBIDDEN"})
	}
}
