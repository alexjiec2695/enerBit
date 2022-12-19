//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"rest_app/app"
	"rest_app/app/config"
	appdi "rest_app/di/app"
	configdi "rest_app/di/config"
	"rest_app/di/databaseservice"
	"rest_app/di/repositories"
	"rest_app/di/server"
	"rest_app/di/usecases"
)

var providerSet = wire.NewSet(
	appdi.ProviderSet,
	server.GinProviderSet,
	usecases.ProviderSet,
	repositories.ProviderSet,
	databaseservice.ProviderSet,
)

// NewApplication wire's injector to create a new `app.Application`
func NewApplication(conf config.AppConfiguration) (*app.Application, error) {
	wire.Build(providerSet)
	return new(app.Application), nil
}

// NewAppConfiguration wire's injector to create a new `config.AppConfiguration`
func NewAppConfiguration() config.AppConfiguration {
	wire.Build(configdi.ProviderSet)
	return config.AppConfiguration{}
}
