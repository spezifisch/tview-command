package context_test

import (
	"testing"

	"github.com/spezifisch/tview-command/keybinding"
	"github.com/stretchr/testify/assert"
)

func TestContextAddLogic(t *testing.T) {
	configPath := "../testdata/TestContextAddLogic.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	queueContext := (*config)["Queue"]
	assert.NotNil(t, queueContext, "Queue context should be present")

	// Test keybindings defined directly in Queue
	assert.Equal(t, "queue.moveTrack", queueContext.Bindings["m"], "Queue-specific binding for key 'm' should be moveTrack")

	// Ensure that bindings from both Default and ListPreset are inherited
	assert.Equal(t, "goToTop", queueContext.Bindings["g"], "Queue context should inherit key 'g' from ListPreset")
	assert.Equal(t, "addToQueue", queueContext.Bindings["a"], "Queue context should inherit key 'a' from Default")

	// Test that keys not defined or inherited do not exist
	_, exists := queueContext.Bindings["x"]
	assert.False(t, exists, "Binding for 'x' should not exist in Queue")

	_, exists = queueContext.Bindings["z"]
	assert.False(t, exists, "Binding for 'z' should not exist in Queue")
}

func TestContextFallback(t *testing.T) {
	configPath := "../testdata/TestContextFallback.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	queueContext := (*config)["Queue"]

	// Ensure that bindings present in the Queue context are resolved
	assert.Equal(t, "queue.moveTrack", queueContext.Bindings["m"], "Binding for key 'm' should be defined in Queue")

	// Check fallback to Default context for undefined bindings in Queue
	defaultContext := (*config)["Default"]
	assert.Equal(t, "deleteTrack", defaultContext.Bindings["d"], "Default binding for key 'd' should be deleteTrack")
	assert.Equal(t, defaultContext.Bindings["d"], queueContext.Bindings["d"], "Queue should inherit 'd' from Default")
}

func TestContextAddInheritance(t *testing.T) {
	configPath := "../testdata/TestContextAddInheritance.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	queueContext := (*config)["Queue"]
	assert.NotNil(t, queueContext, "Queue context should be present")

	// Ensure keybindings from Default are inherited in Queue
	assert.Equal(t, "deleteTrack", queueContext.Bindings["d"], "Binding for key 'd' should be inherited from Default")
	assert.Equal(t, "addToQueue", queueContext.Bindings["a"], "Binding for key 'a' should be inherited from Default")
}
