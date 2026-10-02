// Package grouter teaches Gin routing.
package grouter

import "github.com/gin-gonic/gin"

// NewRouter serves GET /tickets with "list".
// TODO: gin.New() + r.GET("/tickets", ...) with c.String(200, "list").
func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
