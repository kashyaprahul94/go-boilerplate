package api

import (
	"github.com/kashyaprahul94/go-boilerplate/internal/api/health"
	"github.com/kashyaprahul94/go-boilerplate/internal/api/users"

	"github.com/kashyaprahul94/go-boilerplate/internal/http"
)

func InitRoutes(router http.HttpRouter) {

	// Init `health check` routes
	health.InitApiRoutes(router)

	// Init `user` routes
	users.InitApiRoutes(router)
}
