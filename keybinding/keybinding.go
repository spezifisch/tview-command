package keybinding

import (
	"github.com/BurntSushi/toml"

	"github.com/spezifisch/tview-command/context"
)

type Config map[string]context.Context

// LoadConfig loads a config.toml file from path,
// validates the "keybinding graph" and parses it.
func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	//log.Printf("Config: %+v\n", config)

	// Validate the config for cycles and maybe other brokenness
	if err := ValidateConfig(config); err != nil {
		return nil, err
	}

	// Check if config is essentially empty and warn if so
	hasBindings := false
	for _, context := range config {
		if len(context.Bindings) > 0 {
			hasBindings = true
			break
		}
	}

	if !hasBindings {
		logMessage("Warning: Config has no bindings defined.")
	}

	logMessage("Config loaded.")

	// It's okay if the config is empty
	return &config, nil
}
