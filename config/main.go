// Package config provides the global configuration for the service
package config

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MichaelSimkin/go-template/utils"
	"github.com/gobuffalo/envy"
)

// Config holds the global configuration for the service
var Config = &config{}

// config defines the global configuration structure
type config struct {
	Service struct {
		Port string `json:"port"`
	} `json:"service"`
	Mongo struct {
		URI                   string        `json:"URI"`
		FeatureCollectionName string        `json:"featureCollectionName"`
		ConnectionTimeout     time.Duration `json:"connectionTimeout"`
		ClientPingTimeout     time.Duration `json:"clientPingTimeout"`
		CreateIndexTimeout    time.Duration `json:"createIndexTimeout"`
	} `json:"mongo"`
}

// Init loads the configuration from the environment variables
// If values are not set, the default values are used
func Init() {
	loadDotEnv()
	loadEnvVars()
	logPrettyConfig()
}

// loadEnvVars loads the configuration from the environment variables
func loadEnvVars() {
	var errors []error
	var err error

	// Service
	Config.Service.Port = envy.Get("PORT", "3000")

	// Mongo
	Config.Mongo.URI = envy.Get("MONGO_URI", "mongodb://localhost:27017")
	Config.Mongo.FeatureCollectionName = envy.Get("MONGO_FEATURE_COLLECTION_NAME", "features")
	Config.Mongo.ConnectionTimeout, err = time.ParseDuration(envy.Get("MONGO_CONNECTION_TIMEOUT", "10s"))
	if err != nil {
		errors = append(errors, err)
	}
	Config.Mongo.ClientPingTimeout, err = time.ParseDuration(envy.Get("CLIENT_PING_TIMEOUT", "10s"))
	if err != nil {
		errors = append(errors, err)
	}
	Config.Mongo.CreateIndexTimeout, err = time.ParseDuration(envy.Get("CREATE_INDEX_TIMEOUT", "10s"))
	if err != nil {
		errors = append(errors, err)
	}

	// Log all errors
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		log.Fatal("Exiting...")
	}
}

// loadDotEnv loads the .env file if LOAD_DOTENV env variable is truthy
func loadDotEnv() {
	_, ok := utils.Truthy[envy.Get("LOAD_DOTENV", "false")]
	if ok {
		err := envy.Load(envy.Get("DOTENV_FILE_NAME", ".env"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// logPrettyConfig logs the configuration in a pretty format
func logPrettyConfig() {
	prettyConfig, err := json.MarshalIndent(*Config, "    ", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config: %v\n", string(prettyConfig))
}
