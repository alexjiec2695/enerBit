package databaseservice

import (
	"github.com/google/wire"
	"rest_app/database"
	adapter "rest_app/infrastructure/drivenadapters/adapters/enerbit"
	"rest_app/infrastructure/drivenadapters/databaseservice/enerbit"
)

var ProviderSet = wire.NewSet(
	database.NewPostgresStoreConnection,
	enerbit.NewDatabaseServiceImpl,
	database.NewConnectionRedis,
	enerbit.NewRedisServiceImpl,
	wire.Bind(
		new(adapter.AdapterDB),
		new(*enerbit.DatabaseServiceImpl),
	),
	wire.Bind(
		new(adapter.AdapterStreams),
		new(*enerbit.RedisServiceImpl),
	),
)
