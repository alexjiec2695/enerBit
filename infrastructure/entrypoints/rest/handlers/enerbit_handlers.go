package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_app/domain/usecase/enerbit"
)

type Handler struct {
	useCase *enerbit.UseCase
}

func NewHandler(useCase *enerbit.UseCase) *Handler {
	return &Handler{useCase}
}

func (e *Handler) Test(ctx *gin.Context) {
	resultado := e.useCase.TestUsecase()
	ctx.JSON(http.StatusOK, resultado)
}
