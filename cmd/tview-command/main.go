package main

import (
	"log"

	"github.com/spezifisch/tview-command/keybinding"
)

func main() {
	configPath := "example.toml"

	// Load the configuration file
	config, err := keybinding.LoadConfig(configPath)
	if err != nil || config == nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	log.Println("Configuration loaded successfully.")
}
