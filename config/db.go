package config

func getDbConfig(envVars map[string]string) *DBConfig {
	return &DBConfig{
		MySQL: getMySqlConfig(envVars),
	}
}

func getMySqlConfig(envVars map[string]string) *MySqlConfig {
	host := envVars["MYSQL_DB_HOST"]
	port := envVars["MYSQL_DB_PORT"]

	return &MySqlConfig{
		Host: host,
		Port: port,
	}
}
