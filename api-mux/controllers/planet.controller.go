package controllers

import "github.com/gin-gonic/gin"

func FindAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "findAll",
	})
}

func FindByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "findById",
	})
}

func Create(c *gin.Context) {
	c.JSON(201, gin.H{
		"method": "Create",
	})
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "update",
	})
}

func Delete(c *gin.Context) {
	c.Status(204)
}
