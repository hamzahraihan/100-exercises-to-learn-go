// Package gbind teaches Gin JSON binding.
package gbind

import (
	"github.com/gin-gonic/gin"
)

// Ticket is a support request.
type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// NewRouter serves POST /tickets: 201 + JSON on valid title, 400 otherwise.
// TODO: r.POST("/tickets", ...) with c.ShouldBindJSON, Title check,
// c.JSON(http.StatusCreated, t) / c.JSON(http.StatusBadRequest, ...).
func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
