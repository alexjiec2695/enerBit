package config

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const defaultConfigPath string = "environment/"
const configPathEnv string = "CONFIG_PATH"

// InitAppConfiguration is the initial function to load app configuration
func InitAppConfiguration() AppConfiguration {
	return loadAppProperties()
}

func loadAppProperties() AppConfiguration {
	log.Println("Configuring read of the application properties file.")
	err := putViperConfiguration()
	if err != nil {
		log.Panic(fmt.Errorf("Fatal error setting application configuration: %s", err))
	}

	log.Println("Reading application configuration file.")
	err = readConfigFile()
	if err != nil {
		log.Panic(fmt.Errorf("Fatal error reading application configuration file: %s", err))
	}

	appConfiguration := new(AppConfiguration)
	err = viper.Unmarshal(&appConfiguration)
	if err != nil {
		log.Panic(fmt.Errorf("Fatal error unmarshalling application configuration into an instance %s", err))
	}

	log.Println("Application configuration file loaded successfully. :)")

	return *appConfiguration
}

func putViperConfiguration() error {
	path, err := getWorkDirPath()
	if err != nil {
		return err
	}

	configDirectory := fmt.Sprintf("%s/%s", path, defaultConfigPath)

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv(configPathEnv))
	viper.AddConfigPath(".")
	viper.AddConfigPath(defaultConfigPath)
	viper.AddConfigPath(configDirectory)

	return nil
}

func readConfigFile() error {
	var err error = viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}

func getWorkDirPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return path, err
	}

	var operativeSystem string = runtime.GOOS
	switch operativeSystem {
	case "windows":
		return strings.ReplaceAll(path, `\`, "/"), nil
	default:
		return path, nil
	}
}
