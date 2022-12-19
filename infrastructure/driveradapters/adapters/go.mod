module rest_app/infrastructure/drivenadapters/adapters

go 1.15


require (
	rest_app/domain/model v0.0.0
)

replace (
	rest_app/domain/model => ./../../../domain/model
)
