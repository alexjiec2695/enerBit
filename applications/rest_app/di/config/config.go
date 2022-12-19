package config

import (

	"github.com/google/wire"
	"rest_app/app/config"
)

//ProviderSet is the provider configured for wire dependency injection tool
var ProviderSet = wire.NewSet(
	config.InitAppConfiguration,
)
