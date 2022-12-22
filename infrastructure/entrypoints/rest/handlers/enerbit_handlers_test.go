package handlers_test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/usecase/enerbit"
	"rest_app/infrastructure/entrypoints/rest/dtos/rq"
	"rest_app/infrastructure/entrypoints/rest/handlers"
	"strings"
	"testing"
	"time"
)

type GatewayMock struct {
	mock.Mock
}

func (m *GatewayMock) ExistSerialBrand(serial, brand string) (bool, error) {
	args := m.Called(serial, brand)
	return args.Bool(0), args.Error(1)
}

func (m *GatewayMock) ExistMeterInProperty(address string) (bool, error) {
	args := m.Called(address)
	return args.Bool(0), args.Error(1)
}

func (m *GatewayMock) IsActive(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *GatewayMock) Create(data entities.EnerBitEntities) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *GatewayMock) Update(data entities.EnerBitEntities) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *GatewayMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *GatewayMock) GetDisabledMeters() (*[]entities.EnerBitEntities, error) {
	args := m.Called()
	return args.Get(0).(*[]entities.EnerBitEntities), args.Error(1)
}

func (m *GatewayMock) Filter(serial, brand string) (*entities.EnerBitEntities, error) {
	args := m.Called(serial, brand)
	return args.Get(0).(*entities.EnerBitEntities), args.Error(1)
}

type StreamsMock struct {
	mock.Mock
}

func (m *StreamsMock) PublishTicket(data entities.EnerBitEntities) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestHandler_Create_With_StatusOK(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("ExistSerialBrand", data.Serial, data.Brand).Return(false, nil)
	gt.On("ExistMeterInProperty", data.Address).Return(false, nil)
	gt.On("Create", data).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Create(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestHandler_Create_With_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("ExistSerialBrand", data.Serial, data.Brand).Return(false, nil)
	gt.On("ExistMeterInProperty", data.Address).Return(true, nil)
	gt.On("Create", data).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Create(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}

func TestHandler_Create_With_ShouldBindJSON_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("ExistSerialBrand", data.Serial, data.Brand).Return(false, nil)
	gt.On("ExistMeterInProperty", data.Address).Return(true, nil)
	gt.On("Create", data).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	bytes, _ := json.Marshal("{test: 1}")
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Create(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}

func TestHandler_Delete_With_StatusOK(t *testing.T) {
	gt := new(GatewayMock)

	ID := "1234"

	gt.On("IsActive", ID).Return(false, nil)
	gt.On("Delete", ID).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	ginContext.Params = []gin.Param{
		gin.Param{Key: "id", Value: ID},
	}

	handler.Delete(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestHandler_Delete_With_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	ID := "1234"

	gt.On("IsActive", ID).Return(true, nil)
	gt.On("Delete", ID).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	ginContext.Params = []gin.Param{
		gin.Param{Key: "id", Value: ID},
	}

	handler.Delete(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}

func TestHandler_Update_With_StatusOK(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Update", data).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Update(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestHandler_Update_With_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Update", data).Return(errors.New("error"))

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Update(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestHandler_Update_With_ShouldBindJSON_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Update", data).Return(nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	bytes, _ := json.Marshal("{test:1}")
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Update(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestHandler_GetDisabledMeters_With_StatusOK(t *testing.T) {
	gt := new(GatewayMock)

	data := []entities.EnerBitEntities{
		{
			ID:               "1234",
			Brand:            "1234567",
			Address:          "carrera",
			InstallationDate: time.Time{},
			RetirementDate:   time.Time{},
			Serial:           "987654",
			Lines:            1,
			IsActive:         false,
			CreatedAt:        time.Time{},
		},
	}

	gt.On("GetDisabledMeters").Return(&data, nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	handler.GetDisabledMeters(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestHandler_GetDisabledMeters_With_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := []entities.EnerBitEntities{
		{
			ID:               "1234",
			Brand:            "1234567",
			Address:          "carrera",
			InstallationDate: time.Time{},
			RetirementDate:   time.Time{},
			Serial:           "987654",
			Lines:            1,
			IsActive:         false,
			CreatedAt:        time.Time{},
		},
	}

	gt.On("GetDisabledMeters").Return(&data, errors.New("error"))

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	handler.GetDisabledMeters(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestHandler_Filter_With_StatusOK(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Filter", data.Serial, data.Brand).Return(&data, nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Filter(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestHandler_Filter_With_ShouldBindJSON_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Filter", data.Serial, data.Brand).Return(&data, nil)

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	bytes, _ := json.Marshal("{test:1}")
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Filter(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}

func TestHandler_Filter_With_StatusBadRequest(t *testing.T) {
	gt := new(GatewayMock)

	data := entities.EnerBitEntities{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	gt.On("Filter", data.Serial, data.Brand).Return(&data, errors.New("error"))

	uc := enerbit.NewUseCase(gt)
	handler := handlers.NewHandler(uc)

	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	request := rq.Request{
		ID:               "1234",
		Brand:            "1234567",
		Address:          "carrera",
		InstallationDate: time.Time{},
		RetirementDate:   time.Time{},
		Serial:           "987654",
		Lines:            1,
		IsActive:         false,
		CreatedAt:        time.Time{},
	}

	bytes, _ := json.Marshal(request)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	ginContext.Request = &http.Request{
		Body: r,
	}

	handler.Filter(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}
