// Package gmw teaches Gin middleware.
package gmw

import "github.com/gin-gonic/gin"

// WithHeader tags every response with X-Course: go, then continues.
// TODO: c.Header("X-Course", "go") before c.Next().
func WithHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
