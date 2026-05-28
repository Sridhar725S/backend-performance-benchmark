package main

import (
	"benchmark/database"
	"benchmark/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// CONNECT DATABASE
	database.Connect()

	// CREATE GIN ROUTER
	router := gin.Default()

	// LOAD ROUTES
	routes.SetupRoutes(router)

	// START SERVER
	router.Run(":8083")
}