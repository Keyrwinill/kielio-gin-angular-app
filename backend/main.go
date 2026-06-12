package main

import (
	"kielio-gin-angular-app/backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
/*
import (
	"log"

	"kielio-gin-angular-app/backend/internal/config"
	"kielio-gin-angular-app/backend/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(nil)

	r := gin.Default()

	r.Run(":8080")
}
*/