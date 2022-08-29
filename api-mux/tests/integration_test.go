package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GuilhermeRX/API-StarWars/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindAllRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", controllers.FindAll)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())

}

func TestFindByIdRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/:id", controllers.FindByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/2", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())

}

func TestPostRoute(t *testing.T) {
	router := gin.Default()
	router.POST("/", controllers.Create)

	newPlanet := strings.NewReader(`{
		name":    "AnyName",
		terrain: "anyterrain",
		climate: "anyClimate",
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", newPlanet)

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

}

func TestDeleteRoute(t *testing.T) {
	router := gin.Default()
	router.DELETE("/:id", controllers.Delete)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/2", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	// assert.Equal(t, "pong", w.Body.String())

}
