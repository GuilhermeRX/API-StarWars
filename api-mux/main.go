package main

import "github.com/GuilhermeRX/API-StarWars/server"

func main() {
	server := server.NewServer()

	server.Run()
}
