package config

import (
	"fmt"
	"os"
	"strings"
)

// The singleton application configuration data holder
var appConfig AppConfig

// GetAppConfig returns the application config
func GetAppConfig() AppConfig {
	return appConfig
}

func init() {
	// useful to map all the os environment variables to key value pair
	envVars := make(map[string]string)

	// Loop over all the environment variables and map them to key value pair
	for _, envVar := range os.Environ() {
		pair := strings.SplitN(envVar, "=", 2)
		envVars[pair[0]] = pair[1]
	}

	// Set the application environment
	appConfig.Environment = getApplicationEnvironment(envVars)

	// Set the configuration for HTTP server
	appConfig.HttpServer = getHttpServerConfig(envVars)

	// Set the configuration for various persistence services
	appConfig.Persistence = getPersistenceConfig(envVars)
}

// getApplicationEnvironment returns the current application environment
func getApplicationEnvironment(configVars map[string]string) AppEnvironment {
	environment := configVars["APP_ENVIRONMENT"]

	// We really don't want to work with the application instance where we are unaware of the environment
	if environment == "" {
		panicMissingEnvVariable("APP_ENVIRONMENT")
	}

	return AppEnvironment(environment)
}

// panicMissingEnvVariable helps to panic the go application in case a required config variable is missing.
// Helpful because we don't want to work with an application with partial configurations.
func panicMissingEnvVariable(varName string) {
	panic(fmt.Sprintf("missing environment variable -> %s ", varName))
}
