package main

import (
	"log"

	"kielio-gin-angular-app/backend/internal/config"
	"kielio-gin-angular-app/backend/internal/database"
	"kielio-gin-angular-app/backend/internal/handlers"
	"kielio-gin-angular-app/backend/internal/repositories"
	"kielio-gin-angular-app/backend/internal/routes"
	"kielio-gin-angular-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(nil)

	repo := repositories.DeutschAdjektivRepository{
		DB: db,
	}

	service := services.DeutschAdjektivService{
		Repo: repo,
	}

	handler := handlers.DeutschAdjektivHandler{
		Service: service,
	}

	r := gin.Default()

	routes.RegisterRoutes(
		r,
		&handler,
	)

	r.Run(":8080")
}