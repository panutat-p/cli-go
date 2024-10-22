package internal

type Config struct {
	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLPort     string `env:"MYSQL_PORT"`
	MySQLDatabase string `env:"MYSQL_DATABASE"`
	MySQLUsername string `env:"MYSQL_USERNAME"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
}
