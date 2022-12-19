module rest_app/infrastructure/drivenadapters/databaseservice

go 1.15

require (
	github.com/go-resty/resty/v2 v2.3.0
	github.com/jarcoal/httpmock v1.0.6
	github.com/stretchr/testify v1.6.1
	rest_app/domain/model v0.0.0
)

replace rest_app/domain/model => ./../../../domain/model
