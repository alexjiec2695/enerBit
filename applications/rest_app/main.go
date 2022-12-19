package main

import (
	"fmt"
	"rest_app/app/config"
	"rest_app/di"
)

var conf config.AppConfiguration

func init() {
	conf = di.NewAppConfiguration()
}

func main() {
	app, err := di.NewApplication(conf)

	if err != nil {
		panic(fmt.Errorf("Fatal error building the application: %s", err))
	}

	app.Run()
}
