package context

// Context represents a specific context or mode in the application.
type Context struct {
	Name            string
	Bindings        map[string]string
	ContextOverride *string
	ContextAdd      []string
	TextField       bool
	Modal           bool
}

// NewContext creates a new context with given parameters.
func NewContext(name string, bindings map[string]string, contextOverride *string, contextAdd []string, textField bool, modal bool) *Context {
	return &Context{
		Name:            name,
		Bindings:        bindings,
		ContextOverride: contextOverride,
		ContextAdd:      contextAdd,
		TextField:       textField,
		Modal:           modal,
	}
}
