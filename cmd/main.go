package main

import (
	"log"

	"github.com/helloDevAman/movie-base/cmd/app"
	"github.com/helloDevAman/movie-base/config"
)

func main() {

	cfgLoader := &config.EnvConfigLoader{}

	cfg, err := config.LoadNewConfig(cfgLoader)

	if err != nil {
		log.Printf("Unable to load config: %v", err)
	}

	app.Run(cfg)
}
