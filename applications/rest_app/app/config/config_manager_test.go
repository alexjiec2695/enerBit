package config_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"rest_app/app/config"
	"testing"
)

func TestInitAppConfiguration(test *testing.T) {
	os.Setenv("CONFIG_PATH", "../../environment")
	var configApp config.AppConfiguration = config.InitAppConfiguration()

	assert.Equal(test, "dev", configApp.Application.Environment)
	assert.Equal(test, "rest-app-api", configApp.Application.Name)
	assert.NotNil(test, configApp.Server)
}
