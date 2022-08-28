package routes

import (
	"github.com/GuilhermeRX/API-StarWars/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		planets := main.Group("planets")
		{
			planets.GET("/", controllers.FindAll)
			planets.GET("/:id", controllers.FindByID)
			planets.POST("/", controllers.Create)
			planets.DELETE("/:id", controllers.Delete)
		}
	}

	return router
}
