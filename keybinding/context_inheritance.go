package keybinding

import (
	"fmt"

	"github.com/spezifisch/tview-command/context"
)

// resolveContextInheritance merges context_add and context_override chains for the given context.
func resolveContextInheritance(config *Config, contextName string, resolvedContexts map[string]context.Context) error {
	// If this context has already been resolved, return early
	if _, exists := resolvedContexts[contextName]; exists {
		return nil
	}

	// Get the current context
	currentContext, exists := (*config)[contextName]
	if !exists {
		return fmt.Errorf("context %s does not exist", contextName)
	}

	// Start with a fresh context
	resolved := context.Context{
		Bindings: make(map[string]string),
		Settings: currentContext.Settings,
	}

	// Implicitly inherit from Default unless already inherited
	if contextName != "Default" && !contains(currentContext.ContextAdd, "Default") {
		if err := resolveContextInheritance(config, "Default", resolvedContexts); err != nil {
			// There is no Default config section in this file, so skip inheriting that.
		} else {
			// Merge bindings from Default context
			mergeBindings(&resolved, resolvedContexts["Default"])
		}
	}

	// First, resolve any contexts added via context_add
	for _, parentContext := range currentContext.ContextAdd {
		if err := resolveContextInheritance(config, parentContext, resolvedContexts); err != nil {
			return err
		}
		// Merge bindings from parent context
		mergeBindings(&resolved, resolvedContexts[parentContext])
	}

	// Next, handle any context overrides via context_override
	for _, parentContext := range currentContext.ContextOverride {
		if err := resolveContextInheritance(config, parentContext, resolvedContexts); err != nil {
			return err
		}
		// Override bindings from parent context
		overrideBindings(&resolved, resolvedContexts[parentContext])
	}

	// Finally, add/override the current context's own bindings
	overrideBindings(&resolved, currentContext)

	// Store the resolved context
	resolvedContexts[contextName] = resolved

	return nil
}

// mergeBindings adds bindings from the parent context, without overriding existing ones.
func mergeBindings(resolved *context.Context, parent context.Context) {
	for key, action := range parent.Bindings {
		if _, exists := resolved.Bindings[key]; !exists {
			resolved.Bindings[key] = action
		}
	}
}

// overrideBindings overrides or adds the bindings from the parent context to the current one.
func overrideBindings(resolved *context.Context, parent context.Context) {
	for key, action := range parent.Bindings {
		resolved.Bindings[key] = action // This will override existing bindings
	}
}

// Small helper function to check if the given context is already part of the inheritance list.
func contains(contexts []string, context string) bool {
	for _, c := range contexts {
		if c == context {
			return true
		}
	}
	return false
}
