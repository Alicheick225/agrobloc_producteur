package routes

import (
	"agrobloc_producteur/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	//Routes pour les parcelles
	parcelles := r.Group("/api/parcelles")
	{
		parcelles.GET("user/:user_id", controllers.GetParcellesByUser)
		parcelles.POST("/create", controllers.CreateParcelle)
		parcelles.GET("/", controllers.GetAllParcelles)
		parcelles.DELETE("/delete/:id", controllers.DeleteParcelle)
		parcelles.PUT("/update/:id", controllers.UpdateParcelle)
		parcelles.GET("/:id", controllers.GetParcelleByID)

	}

	//Routes pour les types de cultures
	typesCultures := r.Group("/api/types-cultures")
	{
		typesCultures.POST("/create", controllers.CreateTypeCulture)
		typesCultures.GET("/", controllers.GetAllTypesCultures)
		typesCultures.DELETE("/delete/:id", controllers.DeleteTypeCulture)
		typesCultures.PUT("/update/:id", controllers.UpdateTypeCulture)
		typesCultures.GET("/:id", controllers.GetTypeCultureByID)

	}

}
