package keybinding_test

import (
	"testing"

	"github.com/spezifisch/tview-command/keybinding"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	configPath := "../testdata/TestGlobalContext.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")
	assert.NotNil(t, config.Global, "Config should have non-nil contexts")
	assert.NotEmpty(t, config.Global.Bindings, "Config should have contexts")
}

func TestGlobalContext(t *testing.T) {
	configPath := "../testdata/TestGlobalContext.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	expectedBindings := map[string]string{
		"ESC": "closeModal",
	}

	for key, command := range expectedBindings {
		assert.Equal(t, command, config.Global.Bindings[key], "Binding for key %s should be %s", key, command)
	}
}

func TestKeyBindings(t *testing.T) {
	configPath := "../example.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	defaultContext := config.Default
	assert.NotEmpty(t, "Default context should be present")

	assert.Equal(t, "deleteTrack", defaultContext.Bindings["d"], "Binding for key 'd' should be 'deleteTrack'")
	assert.Equal(t, "addToQueue", defaultContext.Bindings["a"], "Binding for key 'a' should be 'addToQueue'")

	queueContext := config.Queue
	assert.NotEmpty(t, "Queue context should be present")

	assert.Equal(t, "queue.deleteTrack", queueContext.Bindings["d"], "Binding for key 'd' in Queue should be 'queue.deleteTrack'")
	assert.Equal(t, "shuffleQueue", queueContext.Bindings["s"], "Binding for key 's' in Queue should be 'shuffleQueue'")
}
