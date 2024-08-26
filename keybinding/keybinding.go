package keybinding

import (
	"errors"
	"log"

	"github.com/BurntSushi/toml"
)

type Config map[string]Context

type Context struct {
	Bindings        map[string]string `toml:"bindings"`
	ContextAdd      []string          `toml:"context_add"`
	ContextOverride []string          `toml:"context_override"` // Change to slice of strings
}

func validateContextAdd(config Config) error {
	for contextName, context := range config {
		for _, add := range context.ContextAdd {
			if _, exists := config[add]; !exists {
				return errors.New("context_add refers to an invalid or missing context: " + add + " in context: " + contextName)
			}
		}
	}
	return nil
}

func validateContextOverride(config Config) error {
	for contextName, context := range config {
		for _, override := range context.ContextOverride {
			if _, exists := config[override]; !exists {
				return errors.New("context_override refers to an invalid or missing context: " + override + " in context: " + contextName)
			}
		}
	}
	return nil
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatalf("Error parsing config: %s", err)
		return nil, err
	}

	// Validate context_add references
	if err := validateContextAdd(config); err != nil {
		return nil, err
	}

	// Validate context_override references
	if err := validateContextOverride(config); err != nil {
		return nil, err
	}

	log.Printf("Config Loaded: %+v\n", config)
	return &config, nil
}
