package types

// Context represents a specific context or mode in the application.
type Context struct {
	Bindings        map[string]string      `toml:"bindings"`
	ContextAdd      []string               `toml:"context_add,omitempty"`
	ContextOverride []string               `toml:"context_override,omitempty"`
	Settings        map[string]interface{} `toml:"settings,omitempty"`
}
