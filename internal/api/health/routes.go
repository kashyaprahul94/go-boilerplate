package health

import "github.com/kashyaprahul94/go-boilerplate/internal/http"

func InitApiRoutes(router http.HttpRouter) {
	router.GET("/", GetAppInfo)
	router.GET("/health", GetAppHealthInfo)
}
