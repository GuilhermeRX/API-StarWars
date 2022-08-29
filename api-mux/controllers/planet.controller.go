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
		var planetsForName, err = services.FindByName(string)

		if len(planetsForName) == 0 || err != nil {
			c.JSON(400, gin.H{
				"message": "No planets found in the search",
			})
			return
		}

		c.JSON(200, planetsForName)
		return
	}
	var planets = services.FindAll()

	c.JSON(200, planets)

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

	row, err := services.Delete(newid)

	if int64(row) == 0 {
		c.JSON(400, gin.H{
			"message": "ID does not exist",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(204)
}
