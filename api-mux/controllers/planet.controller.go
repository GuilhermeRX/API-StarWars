package controllers

import (
	"strconv"

	"github.com/GuilhermeRX/API-StarWars/models"
	"github.com/GuilhermeRX/API-StarWars/services"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var p = services.FindAll()
	c.JSON(200, p)

}

func FindByID(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}
	c.JSON(200, services.FindByID(newid))
}

func Create(c *gin.Context) {
	var planet models.Planet
	c.ShouldBind(&planet)
	c.JSON(201, services.Create(planet))
}

func Update(c *gin.Context) {
	c.JSON(200, services.Update())
}

func Delete(c *gin.Context) {
	services.Delete()
	c.Status(204)
}
