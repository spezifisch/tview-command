package keybinding_test

import (
	"testing"

	"github.com/spezifisch/tview-command/keybinding"
	"github.com/stretchr/testify/assert"
)

func TestDetectCycle_NoCycle(t *testing.T) {
	config := keybinding.Config{
		"Default": {ContextOverride: []string{"Global"}},
		"Global":  {ContextOverride: []string{}},
	}

	err := keybinding.DetectCycleAndValidate(config)
	assert.NoError(t, err, "Config without cycles should not return an error")
}

func TestDetectCycle_SimpleCycle(t *testing.T) {
	config := keybinding.Config{
		"Default": {ContextOverride: []string{"Global"}},
		"Global":  {ContextOverride: []string{"Default"}},
	}

	err := keybinding.DetectCycleAndValidate(config)
	assert.Error(t, err, "Config with a cycle should return an error")
	assert.Contains(t, err.Error(), "cyclic dependency", "Error message should indicate a cycle")
}

func TestDetectCycle_ComplexCycle(t *testing.T) {
	config := keybinding.Config{
		"Default":  {ContextOverride: []string{"ContextA"}},
		"ContextA": {ContextOverride: []string{"ContextB"}},
		"ContextB": {ContextOverride: []string{"Default"}},
	}

	err := keybinding.DetectCycleAndValidate(config)
	assert.Error(t, err, "Config with a complex cycle should return an error")
	assert.Contains(t, err.Error(), "cyclic dependency", "Error message should indicate a cycle")
}

func TestDetectCycle_ContextAddCycle(t *testing.T) {
	config := keybinding.Config{
		"Default":  {ContextAdd: []string{"ContextA"}},
		"ContextA": {ContextAdd: []string{"ContextB"}},
		"ContextB": {ContextAdd: []string{"Default"}},
	}

	err := keybinding.DetectCycleAndValidate(config)
	assert.Error(t, err, "Config with context_add cycles should return an error")
	assert.Contains(t, err.Error(), "cyclic dependency", "Error message should indicate a cycle")
}
