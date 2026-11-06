package controller

import "github.com/gin-gonic/gin"

// Login authenticates an operator and starts a session.
func Login(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// Logout ends the operator session.
func Logout(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
