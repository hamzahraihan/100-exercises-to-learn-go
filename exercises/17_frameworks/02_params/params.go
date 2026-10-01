// Package gparams teaches Gin path params.
package gparams

import "github.com/gin-gonic/gin"

// NewRouter serves GET /tickets/:id with "one:"+id.
// TODO: r.GET("/tickets/:id", ...) with c.Param("id").
func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
