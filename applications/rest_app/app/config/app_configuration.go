package config

import "rest_app/app/config/mappers"

//AppConfiguration struct represent all app configuration, etc
type AppConfiguration struct {
	Application mappers.Application `mapstructure:"application"`
	Server      mappers.Server      `mapstructure:"server"`
}
