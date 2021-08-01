package config

// getHttpServerConfig returns configuration for http server. Panic if the required variables are not present in the configuration
func getHttpServerConfig(configVars map[string]string) *HttpServerConfig {
	host := configVars["HTTP_SERVER_HOST"]

	// It is fine not to have the host, we can default it to `localhost`
	if host == "" {
		host = "localhost"
	}

	port := configVars["HTTP_SERVER_PORT"]

	// Http server can't really work without a port where it needs to listens to for incoming requests
	if port == "" {
		panicMissingEnvVariable("HTTP_SERVER_PORT")
	}

	return &HttpServerConfig{
		Host: host,
		Port: port,
	}
}
