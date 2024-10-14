package types

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Event struct {
	KeyName       string
	Command       string
	IsBound       bool
	OriginalEvent *tcell.EventKey
	Config        *Config
}

// FromEventKey creates a new Event from a tcell.EventKey and sets the config
func FromEventKey(ev *tcell.EventKey, config *Config) *Event {
	e := &Event{
		OriginalEvent: ev,
		Config:        config, // Assign the config
	}

	if ev.Key() == tcell.KeyRune {
		e.KeyName = fmt.Sprintf("Rune[%c]", ev.Rune())
	} else {
		e.KeyName = ev.Name()
	}

	return e
}

func (e *Event) String() string {
	if e.IsBound {
		return fmt.Sprintf("Key: %s, Command: %s", e.KeyName, e.Command)
	}
	return fmt.Sprintf("Key: %s (unbound)", e.KeyName)
}

func (e *Event) LookupCommand(contextKey string) error {
	if e.Config == nil {
		return fmt.Errorf("tviewcommand.types.Event.Config is nil")
	}

	// If the contextKey is in the format "Rune[%s]", strip "Rune[" and "]"
	if strings.HasPrefix(contextKey, "Rune[") && strings.HasSuffix(contextKey, "]") {
		contextKey = strings.TrimSuffix(strings.TrimPrefix(contextKey, "Rune["), "]")
	}

	// Fetch the current context based on the contextKey
	currentContext, ok := (*e.Config)[contextKey]
	if !ok {
		e.IsBound = false
		e.Command = ""
		return fmt.Errorf("Lookup failed: Context '%s' not found.", contextKey)
	}

	// Check if the KeyName from the event has a command bound to it in the current context
	if command, found := currentContext.Bindings[e.KeyName]; found {
		e.Command = command
		e.IsBound = true
	} else {
		e.Command = ""
		e.IsBound = false
	}

	return nil
}
