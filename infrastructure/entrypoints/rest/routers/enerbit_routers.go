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
	group.POST("meters", h.handler.Create)
	group.PUT("meters", h.handler.Update)
	group.DELETE("meters/:id", h.handler.Delete)
	group.GET("meters/disabled", h.handler.GetDisabledMeters)
	group.POST("meters/filter", h.handler.Filter)
}
