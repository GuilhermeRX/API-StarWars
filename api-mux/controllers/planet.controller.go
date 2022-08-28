package controllers

import (
	"strconv"

	"github.com/GuilhermeRX/API-StarWars/models"
	"github.com/GuilhermeRX/API-StarWars/services"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	string, bool := c.GetQuery("name")
	if bool == true {
		var filterName, err = services.FindByName(string)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, filterName)
		return
	}
	var p = services.FindAll()
	c.JSON(200, p)

}

func FindByID(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "ID has to be integer",
		})
		return
	}

	planet, err := services.FindByID(newid)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, planet)
}

func Create(c *gin.Context) {
	var planet models.Planet
	c.ShouldBind(&planet)

	err := services.ValidateStruct(planet)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	planetCreated, err := services.Create(planet)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(201, planetCreated)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	services.Delete(newid)
	c.Status(204)
}
