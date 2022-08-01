package main

import (
	"log"
	"youtube-downloader-rest/config"
	"youtube-downloader-rest/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
