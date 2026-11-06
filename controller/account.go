package controller

import "github.com/gin-gonic/gin"

// ListAccounts returns onboarded accounts for the admin UI.
func ListAccounts(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// UpdateAccount edits account notes / fingerprint binding.
func UpdateAccount(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
