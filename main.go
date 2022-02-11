package main

import (
	"github.com/MichaelSimkin/go-template/config"
	"github.com/MichaelSimkin/go-template/server"
)

func main() {
	config.Init()
	server.NewServer(config.Config.Service.Port).Run()
}
