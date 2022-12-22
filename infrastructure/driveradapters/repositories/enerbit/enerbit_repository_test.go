package enerbit_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/repositories/enerbit"
	"testing"
	"time"
)

type AdapterMock struct {
	mock.Mock
}

func (m *AdapterMock) ExistSerialBrand(serial, brand string) (bool, error) {
	args := m.Called(serial, brand)
	return args.Bool(0), args.Error(1)
}

func (m *AdapterMock) ExistMeterInProperty(address string) (bool, error) {
	args := m.Called(address)
	return args.Bool(0), args.Error(1)
}

func (m *AdapterMock) IsActive(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *AdapterMock) Create(data entities.EnerBitEntities) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *AdapterMock) Update(data entities.EnerBitEntities) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *AdapterMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *AdapterMock) GetDisabledMeters() (*[]entities.EnerBitEntities, error) {
	args := m.Called()
	return args.Get(0).(*[]entities.EnerBitEntities), args.Error(1)
}

func (m *AdapterMock) Filter(serial, brand string) (*entities.EnerBitEntities, error) {
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

func TestRepository_Create_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{}
	am.On("Create", data).Return(nil)
	streams.On("PublishTicket", data).Return(nil)

	r := enerbit.NewRepository(&am, &streams)

	err := r.Create(data)

	assert.NoError(t, err)
}

func TestRepository_Update_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{}
	am.On("Update", data).Return(nil)

	r := enerbit.NewRepository(&am, &streams)

	err := r.Update(data)

	assert.NoError(t, err)
}

func TestRepository_ExistSerialBrand_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{
		Brand:  "987654",
		Serial: "123456",
	}
	am.On("ExistSerialBrand", data.Serial, data.Brand).Return(true, nil)

	r := enerbit.NewRepository(&am, &streams)

	exist, err := r.ExistSerialBrand(data.Serial, data.Brand)

	assert.Equal(t, true, exist)
	assert.NoError(t, err)
}

func TestRepository_ExistMeterInProperty_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{
		Brand:   "987654",
		Serial:  "123456",
		Address: "carrera",
	}
	am.On("ExistMeterInProperty", data.Address).Return(true, nil)

	r := enerbit.NewRepository(&am, &streams)

	exist, err := r.ExistMeterInProperty(data.Address)

	assert.Equal(t, true, exist)
	assert.NoError(t, err)
}

func TestRepository_IsActive_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{
		ID:      "12312361283",
		Brand:   "987654",
		Serial:  "123456",
		Address: "carrera",
	}
	am.On("IsActive", data.ID).Return(true, nil)

	r := enerbit.NewRepository(&am, &streams)

	exist, err := r.IsActive(data.ID)

	assert.Equal(t, true, exist)
	assert.NoError(t, err)
}

func TestRepository_Delete_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
	data := entities.EnerBitEntities{
		ID: "12312361283",
	}
	am.On("Delete", data.ID).Return(nil)

	r := enerbit.NewRepository(&am, &streams)

	err := r.Delete(data.ID)

	assert.NoError(t, err)
}

func TestRepository_GetDisabledMeters_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
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
	am.On("GetDisabledMeters").Return(&data, nil)

	r := enerbit.NewRepository(&am, &streams)

	response, err := r.GetDisabledMeters()

	assert.NoError(t, err)
	assert.Equal(t, data, *response)
}

func TestRepository_Filter_Success(t *testing.T) {
	am := AdapterMock{}
	streams := StreamsMock{}
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

	am.On("Filter", data.Serial, data.Brand).Return(&data, nil)

	r := enerbit.NewRepository(&am, &streams)

	response, err := r.Filter(data.Serial, data.Brand)

	assert.NoError(t, err)
	assert.Equal(t, data, *response)
}
