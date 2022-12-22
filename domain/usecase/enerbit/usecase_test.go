package enerbit_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/usecase/enerbit"
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

func TestUseCase_Create_Successful(t *testing.T) {
	// Arrange
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

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Create(data)
	// Assert
	assert.NoError(t, err)
}

func TestUseCase_Create_Return_Error_ExistSerialBrand(t *testing.T) {
	// Arrange
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

	gt.On("ExistSerialBrand", data.Serial, data.Brand).Return(false, errors.New("error"))
	gt.On("ExistMeterInProperty", data.Address).Return(false, nil)
	gt.On("Create", data).Return(nil)

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Create(data)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "consulting database error")
}

func TestUseCase_Create_Return_Error_ExistMeterInProperty(t *testing.T) {
	// Arrange
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
	gt.On("ExistMeterInProperty", data.Address).Return(false, errors.New("error"))
	gt.On("Create", data).Return(nil)

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Create(data)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "consulting database error")
}

func TestUseCase_Create_Successful_ExistSerialBrand_True(t *testing.T) {
	// Arrange
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

	gt.On("ExistSerialBrand", data.Serial, data.Brand).Return(true, nil)
	gt.On("ExistMeterInProperty", data.Address).Return(false, nil)
	gt.On("Create", data).Return(nil)

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Create(data)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "business error, serial or brand exists")
}

func TestUseCase_Create_Successful_ExistMeterInProperty_True(t *testing.T) {
	// Arrange
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

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Create(data)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "business error, meter active in this address")
}

func TestUseCase_Delete_Successful(t *testing.T) {
	// Arrange
	gt := new(GatewayMock)
	ID := "1234"
	gt.On("IsActive", ID).Return(false, nil)
	gt.On("Delete", ID).Return(nil)

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Delete(ID)
	// Assert
	assert.NoError(t, err)
}

func TestUseCase_Delete_Failed_IsActive_Error(t *testing.T) {
	// Arrange
	gt := new(GatewayMock)
	ID := "1234"
	gt.On("IsActive", ID).Return(false, errors.New("error"))

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Delete(ID)
	// Assert
	assert.Error(t, err)
	assert.Error(t, err, "error")
}

func TestUseCase_Delete_Successful_Is_Active(t *testing.T) {
	// Arrange
	gt := new(GatewayMock)
	ID := "1234"
	gt.On("IsActive", ID).Return(true, nil)
	gt.On("Delete", ID).Return(nil)

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Delete(ID)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "business error, the meter is active")
}

func TestUseCase_Delete_Error(t *testing.T) {
	// Arrange
	gt := new(GatewayMock)
	ID := "1234"
	gt.On("IsActive", ID).Return(false, nil)
	gt.On("Delete", ID).Return(errors.New("error"))

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Delete(ID)
	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "error")
}

func TestUseCase_Update_Successful(t *testing.T) {
	// Arrange
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

	// Act
	uc := enerbit.NewUseCase(gt)
	err := uc.Update(data)
	// Assert
	assert.NoError(t, err)
}

func TestUseCase_GetDisabledMeters_Successful(t *testing.T) {
	// Arrange
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

	// Act
	uc := enerbit.NewUseCase(gt)
	res, err := uc.GetDisabledMeters()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, data, *res)
}

func TestUseCase_Filter_Successful(t *testing.T) {
	// Arrange
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

	// Act
	uc := enerbit.NewUseCase(gt)
	res, err := uc.Filter(data)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, data, *res)
}
