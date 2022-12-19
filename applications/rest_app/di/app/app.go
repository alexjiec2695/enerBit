package app

import (

	"github.com/google/wire"
	"rest_app/app"
)

//ProviderSet is the provider configured for wire dependency injection tool
var ProviderSet = wire.NewSet(
	wire.Struct(new(app.Application), "*"),
)
