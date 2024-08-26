package keybinding

import (
	"fmt"
)

func DetectCycleAndValidate(config Config) error {
	visited := make(map[string]bool)

	var visit func(context string) error
	visit = func(context string) error {
		if visited[context] {
			return fmt.Errorf("cyclic dependency detected")
		}

		visited[context] = true
		defer func() { visited[context] = false }()

		for _, nextContext := range config[context].ContextOverride {
			if err := visit(nextContext); err != nil {
				return err
			}
		}

		for _, nextContext := range config[context].ContextAdd {
			if err := visit(nextContext); err != nil {
				return err
			}
		}

		return nil
	}

	for context := range config {
		if err := visit(context); err != nil {
			return err
		}
	}

	return nil
}
