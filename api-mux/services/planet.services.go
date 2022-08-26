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

func FindByID(id int) models.Planet {
	return models.FindByID(id)
}

func Create(planet models.Planet) models.Planet {
	newPlanet := models.Planet{
		ID:      models.Count() + 1,
		Name:    planet.Name,
		Clima:   planet.Clima,
		Terreno: planet.Terreno,
	}
	return models.Create(newPlanet)
}

func Update() Service {
	return Service{
		Method: "update",
	}
}

func Delete() {
	return
}
