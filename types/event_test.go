package types

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestFromEventKey_RuneKey(t *testing.T) {
	ev := tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone)
	event := FromEventKey(ev, nil)

	assert.Equal(t, "Rune[a]", event.KeyName)
	assert.Equal(t, ev, event.OriginalEvent)
}

func TestFromEventKey_SpecialKey(t *testing.T) {
	ev := tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
	event := FromEventKey(ev, nil)

	assert.Equal(t, "Enter", event.KeyName)
	assert.Equal(t, ev, event.OriginalEvent)
}

func TestString_BoundCommand(t *testing.T) {
	event := &Event{
		KeyName: "Enter",
		Command: "Submit",
		IsBound: true,
	}

	expected := "Key: Enter, Command: Submit"
	assert.Equal(t, expected, event.String())
}

func TestString_UnboundCommand(t *testing.T) {
	event := &Event{
		KeyName: "Enter",
		Command: "",
		IsBound: false,
	}

	expected := "Key: Enter (unbound)"
	assert.Equal(t, expected, event.String())
}
func TestLookupCommand_Bound(t *testing.T) {
	config := Config{
		"Main": Context{
			Bindings: map[string]string{
				"Enter": "Submit",
			},
		},
	}

	event := FromEventKey(tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone), &config)
	err := event.LookupCommand("Main")

	assert.Nil(t, err)
	assert.True(t, event.IsBound)
	assert.Equal(t, "Submit", event.Command)
}

func TestLookupCommand_Unbound(t *testing.T) {
	config := Config{
		"Main": Context{
			Bindings: map[string]string{},
		},
	}

	event := FromEventKey(tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone), &config)
	err := event.LookupCommand("Main")

	assert.Nil(t, err)
	assert.False(t, event.IsBound)
	assert.Equal(t, "", event.Command)
}

func TestLookupCommand_ContextKeyWithRunePattern(t *testing.T) {
	config := Config{
		"a": Context{
			Bindings: map[string]string{
				"Rune[a]": "ActionA",
			},
		},
	}

	event := FromEventKey(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone), &config)
	err := event.LookupCommand("Rune[a]")

	assert.Nil(t, err)
	assert.True(t, event.IsBound)
	assert.Equal(t, "ActionA", event.Command)
}

func TestLookupCommand_ContextKeyNotFound(t *testing.T) {
	config := Config{
		"Main": Context{
			Bindings: map[string]string{
				"Enter": "Submit",
			},
		},
	}

	fakeKey := tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
	event := FromEventKey(fakeKey, &config)
	err := event.LookupCommand("NonExistentContext")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Lookup failed: Context 'NonExistentContext' not found.")
	assert.False(t, event.IsBound)
	assert.Equal(t, "", event.Command)
}
