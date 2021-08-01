package config

// getPersistenceConfig returns configuration for different persistence services used in the application.
// Panic if the required variables are not present in the configuration
func getPersistenceConfig(configVars map[string]string) *PersistenceConfig {
	return &PersistenceConfig{
		// Config for MySQL
		MySQL: getMySqlConfig(configVars),
	}
}

// getMySqlConfig  returns configuration for MySQL
func getMySqlConfig(configVars map[string]string) *MySqlConfig {
	host := configVars["MYSQL_DB_HOST"]

	// MySQL can't be reached out to, without the host
	if host == "" {
		panicMissingEnvVariable("MYSQL_DB_HOST")
	}

	port := configVars["MYSQL_DB_PORT"]

	// MySQL can't be reached out to, without the port
	if port == "" {
		panicMissingEnvVariable("MYSQL_DB_PORT")
	}

	user := configVars["MYSQL_DB_USERNAME"]

	// MySQL can't be reached out to, without the username
	if user == "" {
		panicMissingEnvVariable("MYSQL_DB_USERNAME")
	}

	password := configVars["MYSQL_DB_PASSWORD"]

	// MySQL can't be reached out to, without the password (except the cases where there is no password. Highly risky & unlikely!)
	if password == "" {
		panicMissingEnvVariable("MYSQL_DB_PASSWORD")
	}

	databaseName := configVars["MYSQL_DB_DATABASE"]

	// This can be optional but generally our app is centric to only one DB, hence we should always expect this to be there.
	if databaseName == "" {
		panicMissingEnvVariable("MYSQL_DB_DATABASE")
	}

	return &MySqlConfig{
		Host:         host,
		Port:         port,
		User:         user,
		Password:     password,
		DatabaseName: databaseName,
	}
}
