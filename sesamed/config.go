package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/plants/sesame"
	"net/url"

	"log"
)

type SesamedConfig struct {
	Debug      bool
	Port       int
	StorageURL string
	Storage    sesame.UserStore
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

	if config.StorageURL == "" {
		config.StorageURL = "memory://"
	}

	connection_url, err := url.Parse(config.StorageURL)
	if err != nil {
		log.Print("error")
		log.Fatal(err)
	}
	switch connection_url.Scheme {
	case "memory":
		config.Storage, err = sesame.NewInMemoryStore()
	default:
		panic("I do not know how to connect to \"" + connection_url.Scheme + "\"")
	}

	if err != nil {
		log.Fatal(err)
	}

	// TODO: implement a storage backend if Debug is off

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	parseConfig()
}
