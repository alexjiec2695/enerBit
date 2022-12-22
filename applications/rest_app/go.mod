module rest_app

require (
	github.com/PuerkitoBio/purell v1.2.0 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/google/wire v0.5.0
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/spf13/viper v1.14.0
	github.com/stretchr/testify v1.8.1
	github.com/swaggo/files v1.0.0 // indirect
	github.com/swaggo/gin-swagger v1.5.3 // indirect
	github.com/swaggo/swag v1.8.9 // indirect
	github.com/urfave/cli/v2 v2.23.7 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	gorm.io/driver/postgres v1.4.5
	gorm.io/gorm v1.24.2
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
