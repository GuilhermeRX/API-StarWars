package main

import (
	"github.com/GuilhermeRX/API-StarWars/server"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	server := server.NewServer()

	server.Run()
}
