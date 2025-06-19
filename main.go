package main

import (
	"agrobloc_producteur/config"

	"agrobloc_producteur/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	// Très important : appel de RegisterRoutes
	routes.RegisterRoutes(router)

	router.Run(":8080")

}
