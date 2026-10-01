// Package gerrors teaches Gin error envelopes.
package gerrors

import "github.com/gin-gonic/gin"

// NewRouter serves GET /tickets/:id: 200 JSON for "7", 404 {"error":...} else.
// TODO: c.Param check + c.AbortWithStatusJSON(404, gin.H{"error":...}) with return.
func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
