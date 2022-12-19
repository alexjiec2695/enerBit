module rest_app/domain/usecase

require (
	rest_app/domain/model v0.0.0
	github.com/stretchr/testify v1.6.1
)

replace (
	rest_app/domain/model => ../../domain/model
)

go 1.15
