package gateways

import "rest_app/domain/model/entities/entities"

type Gateway interface {
	ExistSerialBrand(string, string) (bool, error)
	ExistMeterInProperty(string) (bool, error)
	IsActive(string) (bool, error)
	Create(entities.EnerBitEntities) error
	Update(entities.EnerBitEntities) error
	Delete(string) error
	GetDisabledMeters() (*[]entities.EnerBitEntities, error)
	Filter(string, string) (*entities.EnerBitEntities, error)
}
