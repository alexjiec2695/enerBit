module rest_app/infrastructure/drivenadapters/databaseservice

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/uuid v1.3.0
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.24.2 // indirect
	gorm.io/gorm v1.24.2
	rest_app/domain/model v0.0.0
)

replace rest_app/domain/model => ./../../../domain/model
