package config

type MySqlConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	MySQL *MySqlConfig
}

type HttpServerConfig struct {
	Host string
	Port string
}

type AppConfig struct {
	DB *DBConfig
	HttpServer *HttpServerConfig
}
