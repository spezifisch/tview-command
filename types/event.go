package types

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Event struct {
	KeyName       string
	Command       string
	IsBound       bool
	OriginalEvent *tcell.EventKey
}

// FromEventKey creates a new Event from a tcell.EventKey and parses the key name
func FromEventKey(ev *tcell.EventKey) *Event {
	e := &Event{
		OriginalEvent: ev,
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

func (e *Event) LookupCommand(contextKey string, config Config) {
	// Fetch the current context based on the contextKey
	currentContext, ok := config[contextKey]
	if !ok {
		e.IsBound = false
		e.Command = ""
		return
	}

	// Check if the KeyName from the event has a command bound to it in the current context
	if command, found := currentContext.Bindings[e.KeyName]; found {
		e.Command = command
		e.IsBound = true
	} else {
		e.Command = ""
		e.IsBound = false
	}
}
