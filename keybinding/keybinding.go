package keybinding

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config map[string]Context

type Context struct {
	Bindings        map[string]string `toml:"bindings"`
	ContextAdd      []string          `toml:"context_add,omitempty"`
	ContextOverride []string          `toml:"context_override,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	//log.Printf("Config: %+v\n", config)

	// Validate the config for cycles
	if err := DetectCycleAndValidate(config); err != nil {
		return nil, err
	}

	// Ensure that the config has at least one context with bindings
	hasBindings := false
	for _, context := range config {
		if len(context.Bindings) > 0 {
			hasBindings = true
			break
		}
	}

	if !hasBindings {
		log.Println("Warning: Config has no bindings defined.")
	}

	// It's okay if the config is empty
	return &config, nil
}
