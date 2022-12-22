package enerbit

import (
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/adapters/enerbit"
)

type Repository struct {
	adapter enerbit.Adapter
}

func NewRepository(adapter enerbit.Adapter) *Repository {
	return &Repository{adapter}
}

func (r *Repository) Create(data entities.EnerBitEntities) error {
	return r.adapter.Create(data)
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
