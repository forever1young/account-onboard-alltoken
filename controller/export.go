package controller

import "github.com/gin-gonic/gin"

// These three endpoints are the platform's only inbound surface toward
// account-pool. The pool is always the caller; the platform never dials out
// to the pool. See docs/export-api.md for the request/response contract.

// PullAccounts returns a batch of accounts pending delivery, starting at cursor.
func PullAccounts(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// AckDelivery confirms the pool has stored a batch of accounts, so the platform
// can dequeue / clear plaintext per its retention policy.
func AckDelivery(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// ReportStatus accepts periodic account health/usage updates from the pool,
// replacing any platform-side polling of the pool.
func ReportStatus(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
