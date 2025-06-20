package main

import (
	"agrobloc_producteur/config"

	"agrobloc_producteur/routes"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"time"
)

func main() {
	config.InitDB()

	router := gin.Default()

	// 🔓 Configuration CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.252.15:4200"}, // Remplace * par ton front en prod
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Très important : appel de RegisterRoutes
	routes.RegisterRoutes(router)

	router.Run(":8080")

}
