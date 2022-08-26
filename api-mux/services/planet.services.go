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

func Create() interface{} {
	return models.Create()
}

func Update() Service {
	return Service{
		Method: "update",
	}
}

func Delete() {
	return
}
