package enerbit

import "rest_app/domain/model/entities/entities"

type Adapter interface {
	Create(entities.EnerBitEntities) error
	Update(entities.EnerBitEntities) error
	ExistSerialBrand(serial, brand string) (bool, error)
	Delete(string) error
	IsActive(id string) (bool, error)
	ExistMeterInProperty(string) (bool, error)
	GetDisabledMeters() (*[]entities.EnerBitEntities, error)
	Filter(string, string) (*entities.EnerBitEntities, error)
}
