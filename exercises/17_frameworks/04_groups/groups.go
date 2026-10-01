// Package ggroups teaches Gin route groups.
package ggroups

import "github.com/gin-gonic/gin"

// NewRouter serves GET /api/v1/tickets with "list".
// TODO: r.Group("/api/v1") + v1.GET("/tickets", ...).
func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
