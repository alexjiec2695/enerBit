package databaseservice

import (
	"github.com/google/wire"
	adapter "rest_app/infrastructure/drivenadapters/adapters/enerbit"
	"rest_app/infrastructure/drivenadapters/databaseservice/enerbit"
)

var ProviderSet = wire.NewSet(
	enerbit.NewDatabaseServiceImpl,
	wire.Bind(
		new(adapter.Adapter),
		new(*enerbit.DatabaseServiceImpl),
	),
)
