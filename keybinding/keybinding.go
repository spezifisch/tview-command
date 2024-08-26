package keybinding

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
	for name, context := range config.Global.Bindings {
		log.Printf("Context Loaded: %s -> %+v\n", name, context)
	}

	log.Printf("Config Loaded: %+v\n", config)
	return &config, nil
}
