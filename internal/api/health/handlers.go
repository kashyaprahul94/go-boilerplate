package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppInfo prints a default welcome message. Can be used to check live status of the application
func GetAppInfo(c *gin.Context) {
	response := map[string]string{
		"status":  "Okay",
		"message": "Welcome to the application",
	}

	c.JSON(http.StatusOK, response)
}

// GetAppInfo prints a health report of the application. Can be used to health check of the application
func GetAppHealthInfo(c *gin.Context) {
	response := map[string]interface{}{
		"status": "Okay",
		"health": map[string]interface{}{
			"webserver": "Okay",
		},
	}

	c.JSON(http.StatusOK, response)
}
