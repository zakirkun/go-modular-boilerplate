package main

import (
	"flag"
	"go-modular-boilerplate/internal/app"
	"go-modular-boilerplate/internal/pkg/config"
	user "go-modular-boilerplate/modules/users"
	"log"
	"os"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {

	// Load configuration
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}

	// Start the application
	app := app.NewApp(&cfg)

	// register modules
	app.RegisterModule(user.NewModule())

	// initialize the application
	if err := app.Initialize(); err != nil {
		log.Fatalf("Error initializing application : %v", err)
		os.Exit(1)
	}

	// Start the application
	app.Start()
}
