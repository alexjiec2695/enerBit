package routers

import (
	"github.com/gin-gonic/gin"
)

const BasePath = "api"

type Router struct {
	R *gin.Engine
}

func NewRouter(r *gin.Engine, enerBitRouter *EnerBitRouter) *Router {

	enerBitRouter.registry(r)
	return &Router{r}
}
