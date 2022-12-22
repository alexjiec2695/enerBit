package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/usecase/enerbit"
	"rest_app/infrastructure/entrypoints/rest/dtos/rq"
)

type Handler struct {
	useCase *enerbit.UseCase
}

func NewHandler(useCase *enerbit.UseCase) *Handler {
	return &Handler{useCase}
}

type Book struct {
	ID     int     `json:"id,omitempty"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   *uint16 `json:"year"`
}

// Delete
// @Summary Delete
// @Tags Delete
// @Description Delete meters
// @typeDocument get-string-by-int
// @document get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} object{data=object{code=int}}
// @Failure 400 {object} object{error=object{code=int,description=string,detail=object}} "Bad Request <br>Detail value type can change depends on main error"
// @Failure 404 {object} object{error=object{code=int,description=string,detail=object}} "Not Found <br>Detail value type can change depends on main error"
// @Failure 409 {object} object{error=object{code=int,description=string,detail=object}} "Conflict <br>Detail value type can change depends on main error"
// @Failure 500 {object} object{error=object{code=int,description=string}}
// @Router /api/meters/{id} [delete]
func (e *Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := e.useCase.Delete(id)

	if err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func (e *Handler) Create(ctx *gin.Context) {
	var input rq.Request

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := e.useCase.Create(mapperRequestToEntity(input)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (e *Handler) Update(ctx *gin.Context) {
	var input rq.Request

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := e.useCase.Update(mapperRequestToEntity(input)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (e *Handler) GetDisabledMeters(ctx *gin.Context) {

	result, err := e.useCase.GetDisabledMeters()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, *result)
	return
}

func (e *Handler) Filter(ctx *gin.Context) {
	var input rq.Request

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := e.useCase.Filter(mapperRequestToEntity(input))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func mapperRequestToEntity(data rq.Request) entities.EnerBitEntities {
	return entities.EnerBitEntities{
		ID:               data.ID,
		Brand:            data.Brand,
		Address:          data.Address,
		InstallationDate: data.InstallationDate,
		RetirementDate:   data.RetirementDate,
		Serial:           data.Serial,
		Lines:            data.Lines,
		IsActive:         data.IsActive,
		CreatedAt:        data.CreatedAt,
	}
}
