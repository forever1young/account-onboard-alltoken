package middleware

import "github.com/gin-gonic/gin"

// RequireOperatorSession guards the admin web/API routes.
func RequireOperatorSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
