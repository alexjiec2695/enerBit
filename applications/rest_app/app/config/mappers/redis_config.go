package mappers

type Redis struct {
	Host string `mapstructure:"host"`
	Pass string `mapstructure:"pass"`
	DB   int    `mapstructure:"db"`
}
