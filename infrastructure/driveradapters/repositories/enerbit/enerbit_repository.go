package enerbit

import (
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/adapters/enerbit"
)

type Repository struct {
	adapter        enerbit.AdapterDB
	adapterStreams enerbit.AdapterStreams
}

func NewRepository(adapter enerbit.AdapterDB, adapterStreams enerbit.AdapterStreams) *Repository {
	return &Repository{adapter, adapterStreams}
}

func (r *Repository) Create(data entities.EnerBitEntities) error {
	err := r.adapter.Create(data)
	if err != nil {
		return err
	}
	return r.adapterStreams.PublishTicket(data)
}

func (r *Repository) Update(data entities.EnerBitEntities) error {
	return r.adapter.Update(data)
}

func (r *Repository) ExistSerialBrand(serial, brand string) (bool, error) {
	return r.adapter.ExistSerialBrand(serial, brand)
}

func (r *Repository) ExistMeterInProperty(address string) (bool, error) {
	return r.adapter.ExistMeterInProperty(address)
}

func (r *Repository) IsActive(id string) (bool, error) {
	return r.adapter.IsActive(id)
}

func (r *Repository) Delete(id string) error {
	return r.adapter.Delete(id)
}

func (r *Repository) GetDisabledMeters() (*[]entities.EnerBitEntities, error) {
	return r.adapter.GetDisabledMeters()
}

func (r *Repository) Filter(serial, brand string) (*entities.EnerBitEntities, error) {
	return r.adapter.Filter(serial, brand)
}

func (r *Repository) PublishTicket(data entities.EnerBitEntities) error {
	return r.adapterStreams.PublishTicket(data)
}
