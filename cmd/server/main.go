package main

import (
	"flag"
	"os"

	"github.com/opencars/opencars/pkg/adapter"

	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/http"
	"github.com/opencars/opencars/internal/storage"
)

var (
	path = flag.String("config", "./config/opencars.toml", "Path to configuration file")
)

func main() {
	flag.Parse()

	// Load configuration.
	settings, err := config.New(*path)
	if err != nil {
		panic(err)
	}

	// Create database connection.
	db, err := storage.New(settings)
	if err != nil {
		panic(err)
	}

	// Initialise database connection.
	err = storage.Migrate(db)
	if err != nil {
		panic(err)
	}

	// Run web server.
	http.Storage = adapter.New(db)
	http.Run(settings.API.Address(), os.Getenv("REGS_BASE_URL"))
}
