package mappers

// Application `struct` to access the application configuration
type Application struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"env"`
}
