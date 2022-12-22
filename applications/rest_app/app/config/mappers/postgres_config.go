package mappers

type Postgres struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	DbName   string `mapstructure:"dbName"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
}
