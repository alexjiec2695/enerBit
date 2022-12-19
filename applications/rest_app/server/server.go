package server

import (
	"github.com/gin-gonic/gin"
)

// NewServer create a new `gin.Engine`
func NewServer() *gin.Engine {
	router := gin.New()
	return router
}


