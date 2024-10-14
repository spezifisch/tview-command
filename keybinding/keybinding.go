package keybinding

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	tcContext "github.com/spezifisch/tview-command/context"
	"github.com/spezifisch/tview-command/log"
	"github.com/spezifisch/tview-command/types"
)

// LoadConfig loads a config.toml file from path,
// validates the "keybinding graph", and parses it.
func LoadConfig(path string) (*types.Config, error) {
	var config types.Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, fmt.Errorf("toml.DecodeFile failed: %v", err)
	}

	//log.Printf("Config: %+v\n", config)

	// Validate the config for cycles and maybe other brokenness
	if err := ValidateConfig(config); err != nil {
		return nil, fmt.Errorf("ValidateConfig failed: %v", err)
	}

	// Check if config is essentially empty and warn if so
	hasBindings := false
	caser := cases.Title(language.English)
	for _, context := range config {
		for key, action := range context.Bindings {
			// Convert "CTRL-L" (with minus and any case) to "Ctrl+L" (with plus)
			if strings.Contains(key, "-") {
				newKey := strings.ReplaceAll(key, "-", "+")
				newKey = caser.String(strings.ToLower(newKey))
				delete(context.Bindings, key)
				context.Bindings[newKey] = action
			}
		}
		if len(context.Bindings) > 0 {
			hasBindings = true
		}
	}
	if !hasBindings {
		log.LogMessage("Warning: Config has no bindings defined.")
	}

	// Resolve inheritance for all contexts
	resolvedContexts := make(map[string]types.Context)
	for contextName := range config {
		if _, exists := resolvedContexts[contextName]; !exists {
			// Resolve this context, and store it in resolvedContexts
			if err := tcContext.Resolve(&config, contextName, resolvedContexts); err != nil {
				return nil, err
			}
		}
	}

	// Replace the original config with the resolved contexts
	for contextName, resolvedContext := range resolvedContexts {
		config[contextName] = resolvedContext
	}

	log.LogMessage("Config loaded.")
	return &config, nil
}
