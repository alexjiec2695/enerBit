package enerbit

import (
	"rest_app/infrastructure/drivenadapters/adapters/enerbit"
)

type Repository struct {
	adapter enerbit.Adapter
}

func NewRepository(adapter enerbit.Adapter) *Repository {
	return &Repository{adapter}
}

func (repository *Repository) GetTest() string {
	return repository.adapter.GetAllAdapter()
}
