// Package config provides the global configuration for the service
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/MichaelSimkin/go-template/utils"
	"github.com/gobuffalo/envy"
)

// Service contains the configuration for the service.
var Service = service{}

// Mongo contains the configuration for MongoDB.
var Mongo = mongo{}

type service struct {
	Port string `json:"port"`
}

type mongo struct {
	URI                   string        `json:"URI"`
	FeatureCollectionName string        `json:"featureCollectionName"`
	ConnectionTimeout     time.Duration `json:"connectionTimeout"`
	ClientPingTimeout     time.Duration `json:"clientPingTimeout"`
	CreateIndexTimeout    time.Duration `json:"createIndexTimeout"`
}

// config defines the global configuration structure.
type config struct {
	Service service `json:"service"`
	Mongo   mongo   `json:"mongo"`
}

const (
	defaultMongoTimeout = "10s"
)

// Init loads the configuration from the environment variables.
// If values are not set, the default values are used.
func Init() {
	loadDotEnv()
	loadEnvVars()
	logPrettyConfig(&config{Service, Mongo})
}

// loadEnvVars loads the configuration from the environment variables.
func loadEnvVars() {
	var errs []error
	var err error

	// Service
	Service.Port = envy.Get("PORT", "3000")

	// Mongo
	Mongo.URI = envy.Get("MONGO_URI", "mongodb://localhost:27017")
	Mongo.FeatureCollectionName = envy.Get("MONGO_FEATURE_COLLECTION_NAME", "features")
	Mongo.ConnectionTimeout, err = time.ParseDuration(envy.Get("MONGO_CONNECTION_TIMEOUT", defaultMongoTimeout))
	if err != nil {
		errs = append(errs, err)
	}
	Mongo.ClientPingTimeout, err = time.ParseDuration(envy.Get("CLIENT_PING_TIMEOUT", defaultMongoTimeout))
	if err != nil {
		errs = append(errs, err)
	}
	Mongo.CreateIndexTimeout, err = time.ParseDuration(envy.Get("CREATE_INDEX_TIMEOUT", defaultMongoTimeout))
	if err != nil {
		errs = append(errs, err)
	}

	logConfigErrors(errs)
}

func logConfigErrors(errs []error) {
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic(errors.New("Error loading configuration"))
	}
}

// loadDotEnv loads the .env file if LOAD_DOTENV env variable is truthy.
func loadDotEnv() {
	_, ok := utils.Truthy[envy.Get("LOAD_DOTENV", "false")]
	if ok {
		err := envy.Load(envy.Get("DOTENV_FILE_NAME", ".env"))
		if err != nil {
			panic(err)
		}
	}
}

// logPrettyConfig logs the configuration in a pretty format.
func logPrettyConfig(config *config) {
	prettyConfig, err := json.MarshalIndent(config, "    ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %v\n", string(prettyConfig))
}
