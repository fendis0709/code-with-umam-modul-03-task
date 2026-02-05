package config

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	DBConn  string `mapstructure:"DB_CONN"`
}
