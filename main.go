package main

import (
	"flag"
	"log"

	"github.com/rgonzalezNetel/rConfig/internal/rconfig"
)

func main() {
	configPath := flag.String("config", "config.toml", "Path to the config file")
	flag.Parse()

	if err := rconfig.GenerateProjectStructure(*configPath); err != nil {
		log.Fatalf("Failed to generate project structure: %v", err)
	}
	log.Println("Project structure generated successfully.")
}
