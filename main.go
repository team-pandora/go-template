package main

import (
	"github.com/MichaelSimkin/go-template/config"
	"github.com/MichaelSimkin/go-template/database"
	"github.com/MichaelSimkin/go-template/server"
)

func main() {
	config.Init()
	database.InitMongo()
	server.Serve(config.Service.Port)
}

// TODO:
// - add tests
// - add CI/CD
// - add authentication
// - add swagger
