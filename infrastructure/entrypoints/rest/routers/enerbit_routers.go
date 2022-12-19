package routers

import (
	"github.com/gin-gonic/gin"
	"rest_app/infrastructure/entrypoints/rest/handlers"
)

type EnerBitRouter struct {
	handler *handlers.Handler
}

func NewEnerBitRouter(handler *handlers.Handler) *EnerBitRouter {
	return &EnerBitRouter{handler}
}

func (h *EnerBitRouter) registry(router *gin.Engine) {
	group := router.Group(BasePath)
	group.GET("test", h.handler.Test)
}
