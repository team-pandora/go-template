package main_test

import (
	"testing"

	"github.com/MichaelSimkin/go-template/config"
	"github.com/MichaelSimkin/go-template/database"
	"github.com/MichaelSimkin/go-template/server"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProject(t *testing.T) {
	config.Init()
	database.InitMongo()
	go server.Serve(config.Service.Port)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feature Test Suite")
}
