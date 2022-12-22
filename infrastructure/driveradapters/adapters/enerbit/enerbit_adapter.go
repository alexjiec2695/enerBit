package enerbit

import "rest_app/domain/model/entities/entities"

type AdapterDB interface {
	Create(entities.EnerBitEntities) error
	Update(entities.EnerBitEntities) error
	ExistSerialBrand(serial, brand string) (bool, error)
	Delete(string) error
	IsActive(id string) (bool, error)
	ExistMeterInProperty(string) (bool, error)
	GetDisabledMeters() (*[]entities.EnerBitEntities, error)
	Filter(string, string) (*entities.EnerBitEntities, error)
}

type AdapterStreams interface {
	PublishTicket(data entities.EnerBitEntities) error
}
