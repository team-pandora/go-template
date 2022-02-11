// Package config provides the global configuration for the service
package config

import (
	"encoding/json"
	"log"

	"github.com/MichaelSimkin/go-template/utils"
	"github.com/gobuffalo/envy"
)

// Config holds the global configuration for the service
var Config *ConfigStruct = &ConfigStruct{}

// ConfigStruct defines the global configuration structure
type ConfigStruct struct {
	Service struct {
		Port string `json:"port"`
	} `json:"service"`
	Mongo struct {
		URI                   string `json:"URI"`
		FeatureCollectionName string `json:"featureCollectionName"`
	} `json:"mongo"`
}

// Init loads the configuration from the environment variables
// If values are not set, the default values are used
func Init() {
	loadEnvVars()
	loadDotEnv()
	logPrettyConfig()
}

// loadEnvVars loads the configuration from the environment variables
func loadEnvVars() {
	Config.Service.Port = envy.Get("PORT", "3000")
	Config.Mongo.URI = envy.Get("MONGO_URI", "mongodb://localhost:27017")
	Config.Mongo.FeatureCollectionName = envy.Get("MONGO_FEATURE_COLLECTION_NAME", "features")
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
