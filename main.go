package main

import (
	"github.com/MichaelSimkin/go-template/config"
	"github.com/MichaelSimkin/go-template/server"
)

func main() {
	config.Init()
	server.Serve(config.Config.Service.Port)
}
