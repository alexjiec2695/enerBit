module rest_app

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/wire v0.4.0
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.8.0
	gorm.io/driver/postgres v1.4.5 // indirect
	gorm.io/gorm v1.24.2 // indirect
	rest_app/domain/model v0.0.0
	rest_app/domain/usecase v0.0.0
	rest_app/infrastructure/drivenadapters/adapters v0.0.0
	rest_app/infrastructure/drivenadapters/databaseservice v0.0.0
	rest_app/infrastructure/drivenadapters/repositories v0.0.0
	rest_app/infrastructure/entrypoints/rest v0.0.0
)

replace (
	rest_app/domain/model => ../../domain/model
	rest_app/domain/usecase => ../../domain/usecase
	rest_app/infrastructure/drivenadapters/adapters => ./../../infrastructure/driveradapters/adapters
	rest_app/infrastructure/drivenadapters/databaseservice => ../../infrastructure/driveradapters/databaseservice
	rest_app/infrastructure/drivenadapters/repositories => ./../../infrastructure/driveradapters/repositories
	rest_app/infrastructure/entrypoints/rest => ../../infrastructure/entrypoints/rest
)

go 1.15
