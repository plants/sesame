package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/plants/sesame"

	"log"
)

type SesamedConfig struct {
	Debug       bool
	Port        int
	StorageType string
	Storage     sesame.UserStore
}

var config SesamedConfig

func parseConfig() {
	err := envconfig.Process("sesamed", &config)
	if err != nil {
		log.Fatal(err)
	}

	// set defaults
	if config.Port == 0 {
		config.Port = 2884
	}
	if config.StorageType == "" {
		config.StorageType = "memory"
	}

	config.Storage, err = sesame.NewInMemoryStore()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	parseConfig()
}
