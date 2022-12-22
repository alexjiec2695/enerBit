package repositories

import (
	"github.com/google/wire"
	"rest_app/domain/model/entities/gateways"
	"rest_app/infrastructure/drivenadapters/repositories/enerbit"
)

var ProviderSet = wire.NewSet(
	enerbit.NewRepository,
	wire.Bind(
		new(gateways.Gateway),
		new(*enerbit.Repository),
	),

	wire.Bind(
		new(gateways.Streams),
		new(*enerbit.Repository),
	),
)
