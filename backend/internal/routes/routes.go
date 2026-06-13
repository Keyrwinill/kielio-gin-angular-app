package routes

import (
	"kielio-gin-angular-app/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	deutschHandler *handlers.DeutschAdjektivHandler,
) {

	api := r.Group("/api")

	RegisterDeutschAdjektivRoutes(
		api,
		deutschHandler,
	)
	/*
		api.GET("/language", handlers.GetLanguages)
		api.GET("/language/:name", handlers.GetByName)
		api.GET("/deutsch-adjektiv", deutschHandler.GetDeutschAdjektiv)
	*/
}

func RegisterDeutschAdjektivRoutes(
	api *gin.RouterGroup,
	handler *handlers.DeutschAdjektivHandler,
) {
	api.GET("/deutsch_adjektiv", handler.GetDeutschAdjektiv)
}
