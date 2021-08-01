package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kashyaprahul94/go-boilerplate/config"
)

// Router type
type HttpRouter = *gin.Engine

// HttpServer is the instance for web server
type HttpServer struct {
	// host of the webserver
	host string

	// port on which web server will listen to
	port string

	// router instance to provide routing
	router HttpRouter
}

// NewHttpRouter returns the instance of the default router
func NewHttpRouter() HttpRouter {
	router := gin.Default()

	return router
}

// NewHttpRouterWithMiddlewares returns the instance of router which has standard middlewares attached to it
func NewHttpRouterWithMiddlewares() HttpRouter {
	router := NewHttpRouter()

	return router
}

// NewHttpServer creates a bare minimum http server
func NewHttpServer(config config.AppConfig) *HttpServer {
	httpServerConfig := config.HttpServer

	if config.Environment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return &HttpServer{
		host:   httpServerConfig.Host,
		port:   httpServerConfig.Port,
		router: NewHttpRouterWithMiddlewares(),
	}
}
