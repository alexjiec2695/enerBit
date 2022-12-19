module rest_app/infrastructure/drivenadapters/databaseservice

go 1.15

require (
	github.com/go-resty/resty/v2 v2.3.0
	github.com/jarcoal/httpmock v1.0.6
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.6.1
	gorm.io/gorm v1.24.2 // indirect
	rest_app/domain/model v0.0.0
)

replace rest_app/domain/model => ./../../../domain/model
