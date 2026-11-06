package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireExportAuth guards /api/export/*: only account-pool may call these
// routes, authenticated by a dedicated API key plus an IP allowlist. This is
// the only inbound surface the platform exposes to the pool; the platform
// itself never initiates a connection to the pool.
func RequireExportAuth(apiKey string, allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if keyMatches(c.GetHeader("X-Export-Api-Key"), apiKey) {
			if ipAllowed(c.ClientIP(), allowedIPs) {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func keyMatches(provided, expected string) bool {
	return provided != "" && subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) == 1
}

func ipAllowed(ip string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	for _, a := range allowed {
		if strings.TrimSpace(a) == ip {
			return true
		}
	}
	return false
}
