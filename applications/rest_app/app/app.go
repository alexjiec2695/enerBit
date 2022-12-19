package app

import (
	"rest_app/app/config"
	"rest_app/infrastructure/entrypoints/rest/routers"
)

// Application struct for router and config
type Application struct {
	Router *routers.Router
	Config config.AppConfiguration
}

// Run method for statrt the app
func (a *Application) Run() {
	a.Router.R.Run(a.Config.Server.Port)
}
