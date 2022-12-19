package server

import (
	"github.com/google/wire"
	"rest_app/infrastructure/entrypoints/rest/handlers"
	"rest_app/infrastructure/entrypoints/rest/routers"
	"rest_app/server"
)

var GinProviderSet = wire.NewSet(
	server.NewServer,
	routers.NewRouter,
	routers.NewEnerBitRouter,
	handlers.NewHandler,
)
