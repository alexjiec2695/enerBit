package mappers

// Server struct for data server
type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
