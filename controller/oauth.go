package controller

import "github.com/gin-gonic/gin"

// GenerateAuthLink returns a Claude.ai authorization URL bound to a fingerprint template.
func GenerateAuthLink(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// HandleCallback accepts the operator-submitted state#code and exchanges it for tokens.
func HandleCallback(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
