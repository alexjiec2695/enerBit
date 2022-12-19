module rest_app/infrastructure/drivenadapters/repositories

go 1.15

require (
	github.com/stretchr/testify v1.6.1
	rest_app/domain/model v0.0.0
	rest_app/infrastructure/drivenadapters/adapters v0.0.0
)

replace (
	rest_app/domain/model => ./../../../domain/model
	rest_app/infrastructure/drivenadapters/adapters => ./../adapters
)
