package logger

import (
	"os"

	"github.com/kashyaprahul94/go-boilerplate/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	appConfig := config.GetAppConfig()

	logger = logrus.New()

	// Log to standard output
	logger.SetOutput(os.Stdout)

	// Set the default log level based on the environment
	if appConfig.Environment.IsProduction() {
		logger.SetLevel(logrus.ErrorLevel)
	} else {
		logger.SetLevel(logrus.TraceLevel)
	}

	// Set the default log format level based on the environment
	if appConfig.Environment.IsProduction() {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
}
