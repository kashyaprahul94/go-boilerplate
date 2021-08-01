package app

import (
	_ "github.com/kashyaprahul94/go-boilerplate/config"
	_ "github.com/kashyaprahul94/go-boilerplate/internal/logger"
	_ "github.com/kashyaprahul94/go-boilerplate/internal/persistence"

	"github.com/kashyaprahul94/go-boilerplate/config"
	"github.com/kashyaprahul94/go-boilerplate/internal/api"
	"github.com/kashyaprahul94/go-boilerplate/internal/http"
	"github.com/kashyaprahul94/go-boilerplate/internal/logger"
)

func Boot() {
	// Get the application config
	appConfig := config.GetAppConfig()

	// Create a new instance of http server
	httpServer := http.NewHttpServer(appConfig)

	// Init API routes
	api.InitRoutes(httpServer.GetRouter())

	// Boot the http server
	httpServer.Boot(func() {
		logger.Info("Application has been booted!")
	})
}
