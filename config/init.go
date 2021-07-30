package config

var appConfig AppConfig

func GetAppConfig() AppConfig {
	return appConfig
}

func init () {
	// Read these values from actual env variables or a file
	someEnvVars := map[string]string{
		"MYSQL_DB_HOST": "localhost",
		"MYSQL_DB_PORT": "33006",
		"HTTP_SERVER_HOST": "",
		"HTTP_SERVER_PORT": "8080",
	}

	dbConfig := getDbConfig(someEnvVars)
	httpServerConfig := getHttpServerConfig(someEnvVars)

	appConfig.DB = dbConfig
	appConfig.HttpServer = httpServerConfig
}