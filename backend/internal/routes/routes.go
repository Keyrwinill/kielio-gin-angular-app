package routes

import (
	"kielio-gin-angular-app/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/language", handlers.GetLanguages)
		api.GET("/language/:name", handlers.GetByName)
	}
}
