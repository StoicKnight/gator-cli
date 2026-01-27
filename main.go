package main

import (
	"fmt"
	"log"

	"github.com/StoicKnight/gator-cli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading the config file: %v", err)
	}
	fmt.Println("Config:", cfg)

	if err := cfg.SetUser("alexis"); err != nil {
		log.Fatalf("Error setting current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading the config file: %v", err)
	}
	fmt.Println("Config 2:", cfg)
}
