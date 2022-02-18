package main

import (
	"github.com/team-pandora/go-template/config"
	"github.com/team-pandora/go-template/database"
	"github.com/team-pandora/go-template/server"
)

func main() {
	config.Init()
	database.InitMongo()
	server.Serve(config.Service.Port)
}
