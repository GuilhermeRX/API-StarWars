package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GuilhermeRX/API-StarWars/models"
	"github.com/asaskevich/govalidator"
)

type SwapiPlanet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

type ResponseAPI struct {
	Count    int           `json:"count"`
	Next     int           `json:"next"`
	Previous int           `json:"previous"`
	Results  []SwapiPlanet `json:"results"`
}

func ValidateStruct(planet models.Planet) error {

	_, err := govalidator.ValidateStruct(planet)
	if err != nil {
		return errors.New("Data was sent incorrectly")
	}
	return nil
}

func FindAll() []models.Planet {
	return models.FindAll()
}

func FindByID(id int) (models.Planet, error) {
	return models.FindByID(id)
}

func FindByName(name string) ([]models.Planet, error) {
	return models.FindByName(name)
}

func GetFilmsNumber(planetName string) int {
	url := "https://swapi.dev/api/planets/?search=" + strings.ReplaceAll(planetName, " ", "%20")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	var swapiPlanet ResponseAPI
	json.Unmarshal(responseData, &swapiPlanet)

	var filmsNumber = len(swapiPlanet.Results[0].FilmURLs)
	return filmsNumber
}

func Create(planet models.Planet) (models.Planet, error) {
	fmt.Print(GetFilmsNumber(planet.Name))
	newPlanet := models.Planet{
		ID:               models.Count() + 1,
		Name:             planet.Name,
		Climate:          planet.Climate,
		Terrain:          planet.Terrain,
		MovieAppearances: GetFilmsNumber(planet.Name),
	}

	return models.Create(newPlanet)
}

func Delete(id int) (int64, error) {
	return models.Delete(id)
}
