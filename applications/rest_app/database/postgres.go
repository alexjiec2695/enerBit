package database

import (
	"fmt"
	"rest_app/app/config"
	"rest_app/infrastructure/drivenadapters/databaseservice/entitydata"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ConfigurationDb struct {
	Host         string
	User         string
	Password     string
	DbName       string
	Port         string
	Schema       string
	MaxIdleConns int
	MaxOpenConns int
}

func NewPostgresStoreConnection(config config.AppConfiguration) (*gorm.DB, error) {
	configuration := ConfigurationDb{
		Host:     config.Postgres.Host,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		DbName:   config.Postgres.DbName,
		Port:     config.Postgres.Port,
		Schema:   config.Postgres.Schema,
	}

	var prefix string

	if configuration.Schema != "" {
		prefix = configuration.Schema + "."
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		configuration.Host,
		configuration.User,
		configuration.Password,
		configuration.DbName,
		configuration.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
		},
	})

	if err != nil {
		return &gorm.DB{}, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return &gorm.DB{}, err
	}

	if configuration.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(configuration.MaxIdleConns)
	}

	if configuration.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(configuration.MaxOpenConns)
	}

	if err := db.AutoMigrate(&entitydata.EnerBitData{}); err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
