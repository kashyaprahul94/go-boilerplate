package config

// AppEnvironment represents the type of env where tha application is running
type AppEnvironment string

const (
	Development AppEnvironment = "development"
	Production  AppEnvironment = "production"
)

// IsProduction checks whether the running environment is production or not
func (e AppEnvironment) IsProduction() bool {
	return e == Production
}

// MySqlConfig type represents the values needed to configure MySQL connection
type MySqlConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

// PersistenceConfig type represents the configuration for different sort of persistence services used in the application,
// for example: MySQL, Redis, MongoDB, etc.
type PersistenceConfig struct {
	MySQL *MySqlConfig
}

// HttpServerConfig represents the values needed to configure a HTTP web server
type HttpServerConfig struct {
	Host string
	Port string
}

// AppConfig type represents the state of configuration needed to empower & configure the whole application
type AppConfig struct {
	// The environment of the application
	Environment AppEnvironment

	// HTTP server configuration
	HttpServer *HttpServerConfig

	// Persistence configurations
	Persistence *PersistenceConfig
}
