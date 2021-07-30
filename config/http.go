package config

func getHttpServerConfig(envVars map[string]string) *HttpServerConfig {
	host := envVars["HTTP_SERVER_HOST"]
	port := envVars["HTTP_SERVER_PORT"]

	return &HttpServerConfig{
		Host: host,
		Port: port,
	}
}
