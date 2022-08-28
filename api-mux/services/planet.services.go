package services

import (
	"github.com/GuilhermeRX/API-StarWars/models"
)

type Service struct {
	Method string
}

func FindAll() []models.Planet {
	return models.FindAll()
}

func FindByID(id int) (models.Planet, error) {
	return models.FindByID(id)
}

func FindByName(name string) (models.Planet, error) {
	return models.FindByName(name)
}

func Create(planet models.Planet) (models.Planet, error) {
	newPlanet := models.Planet{
		ID:      models.Count() + 1,
		Name:    planet.Name,
		Clima:   planet.Clima,
		Terreno: planet.Terreno,
	}
	return models.Create(newPlanet)
}

func Delete(id int) {
	models.Delete(id)
}
