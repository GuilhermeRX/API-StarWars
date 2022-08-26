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
			planets.GET("/", controllers.ShowBook)
		}
	}

	return router
}
