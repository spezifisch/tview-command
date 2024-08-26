package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Global  Context
	Default Context
	Queue   Context
}

type Context struct {
	Bindings        map[string]string `toml:"bindings"`
	ContextAdd      string            `toml:"context_add"`
	ContextOverride string            `toml:"context_override"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatalf("Error parsing config: %s", err)
		return nil, err
	}

	// Log each context to confirm it's being loaded
	log.Printf("Global Context Bindings Loaded: %+v\n", config.Global.Bindings)
	log.Printf("Default Context Bindings Loaded: %+v\n", config.Default.Bindings)
	log.Printf("Queue Context Bindings Loaded: %+v\n", config.Queue.Bindings)

	return &config, nil
}

func main() {
	configPath := "example.toml"

	// Load the configuration file
	config, err := LoadConfig(configPath)
	if err != nil || config == nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	log.Println("Configuration loaded successfully.")
}
