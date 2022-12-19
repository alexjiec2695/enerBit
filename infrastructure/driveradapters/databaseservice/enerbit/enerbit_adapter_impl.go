package enerbit

import "gorm.io/gorm"

type DatabaseServiceImpl struct {
	db *gorm.DB
}

func NewDatabaseServiceImpl(db *gorm.DB) *DatabaseServiceImpl {
	return &DatabaseServiceImpl{db: db}
}

func (e *DatabaseServiceImpl) GetAllAdapter() string {
	return "alexander"
}
