package controllers

import (
	"github.com/GuilhermeRX/API-StarWars/services"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var p = services.FindAll()
	c.JSON(200, p)

}

func FindByID(c *gin.Context) {
	c.JSON(200, services.FindByID())
}

func Create(c *gin.Context) {
	c.JSON(201, services.Create())
}

func Update(c *gin.Context) {
	c.JSON(200, services.Update())
}

func Delete(c *gin.Context) {
	services.Delete()
	c.Status(204)
}
