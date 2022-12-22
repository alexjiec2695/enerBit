package enerbit

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/databaseservice/entitydata"
	"strings"
	"time"
)

type DatabaseServiceImpl struct {
	db *gorm.DB
}

func NewDatabaseServiceImpl(db *gorm.DB) *DatabaseServiceImpl {
	return &DatabaseServiceImpl{db: db}
}

func (e *DatabaseServiceImpl) Create(data entities.EnerBitEntities) error {
	result := e.db.Create(mapperToEntityData(data))

	if result.Error != nil {
		return fmt.Errorf("creating data %s", result.Error.Error())
	}

	return nil
}

func (e *DatabaseServiceImpl) Update(data entities.EnerBitEntities) error {

	result := e.db.Model(&entitydata.EnerBitData{}).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"address":         data.Address,
		"retirement_date": data.RetirementDate,
		"lines":           data.Lines,
		"is_active":       data.IsActive,
		"update_at":       time.Now(),
	})

	if result.Error != nil {
		return fmt.Errorf("updating data %s", result.Error.Error())
	}

	return nil
}

func (e *DatabaseServiceImpl) ExistSerialBrand(serial, brand string) (bool, error) {
	var result []entitydata.EnerBitData
	if err := e.db.Where("serial = ? and brand = ?", serial, brand).Take(&result).Error; err != nil {
		if !strings.EqualFold(err.Error(), "record not found") {
			return false, err
		}
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (e *DatabaseServiceImpl) ExistMeterInProperty(address string) (bool, error) {
	var result []entitydata.EnerBitData
	if err := e.db.Where("address = ? and is_active = ?", address, true).Take(&result).Error; err != nil {
		if !strings.EqualFold(err.Error(), "record not found") {
			return false, err
		}
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (e *DatabaseServiceImpl) IsActive(id string) (bool, error) {
	var result []entitydata.EnerBitData
	if err := e.db.Where("id = ? and is_active = ?", id, true).Take(&result).Error; err != nil {
		if !strings.EqualFold(err.Error(), "record not found") {
			return false, err
		}
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (e *DatabaseServiceImpl) Delete(id string) error {
	var items entitydata.EnerBitData
	if err := e.db.Where("id=?", id).Delete(&items).Error; err != nil {
		return err
	}
	return nil
}

func (e *DatabaseServiceImpl) GetDisabledMeters() (*[]entities.EnerBitEntities, error) {
	var result []entitydata.EnerBitData
	if err := e.db.Where("is_active = ?", false).Take(&result).Error; err != nil {
		if !strings.EqualFold(err.Error(), "record not found") {
			return nil, err
		}
	}

	return mappersEntityDataToEntity(result), nil
}

func (e *DatabaseServiceImpl) Filter(serial, brand string) (*entities.EnerBitEntities, error) {
	var result entitydata.EnerBitData
	if err := e.db.Where("serial = ? and brand = ?", serial, brand).Take(&result).Error; err != nil {
		if strings.EqualFold(err.Error(), "record not found") {
			return nil, nil
		} else {
			return nil, err
		}
	}

	meter := mapperEntityDataToEntity(result)
	return &meter, nil
}

func mapperToEntityData(data entities.EnerBitEntities) *entitydata.EnerBitData {
	return &entitydata.EnerBitData{
		ID:               uuid.New(),
		Brand:            data.Brand,
		Address:          data.Address,
		InstallationDate: data.InstallationDate,
		RetirementDate:   sql.NullTime{Time: data.RetirementDate, Valid: true},
		Serial:           data.Serial,
		Lines:            data.Lines,
		IsActive:         data.IsActive,
		CreatedAt:        time.Now(),
	}
}

func mappersEntityDataToEntity(input []entitydata.EnerBitData) *[]entities.EnerBitEntities {
	result := make([]entities.EnerBitEntities, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = mapperEntityDataToEntity(input[i])
	}
	return &result
}

func mapperEntityDataToEntity(input entitydata.EnerBitData) entities.EnerBitEntities {
	result := entities.EnerBitEntities{}
	result.ID = input.ID.String()
	result.Brand = input.Brand
	result.Address = input.Address
	result.InstallationDate = input.InstallationDate
	result.RetirementDate = input.RetirementDate.Time
	result.Serial = input.Serial
	result.Lines = input.Lines
	result.IsActive = input.IsActive
	result.CreatedAt = input.CreatedAt
	return result
}
