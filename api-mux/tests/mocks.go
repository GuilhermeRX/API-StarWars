package tests

import (
	"github.com/gin-gonic/gin"
)

var AllPlanets = `{
	"id":               1,
	"name":             "AnyName",
	"climate":          "AnyClimate",
	"terrain":          "AnyTerrain",
	"movieAppearances": 2,
},
	{
		"id":               2,
		"name":             "Alderaan",
		"climate":          "AnyClimate",
		"terrain":          "AnyTerrain",
		"movieAppearances": 2,
	},
	{
		"id":               3,
		"name":             "AnyName",
		"climate":          "AnyClimate",
		"terrain":          "AnyTerrain",
		"movieAppearances": 2,
	}}`

func FindAllMock(c *gin.Context) {
	c.JSON(200, AllPlanets)
}
