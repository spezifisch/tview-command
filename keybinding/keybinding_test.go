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
	assert.NotNil(t, (*config)["Global"], "Config should have Global context")
	assert.NotEmpty(t, (*config)["Global"].Bindings, "Global context should have bindings")
}

func TestGlobalContext(t *testing.T) {
	configPath := "../testdata/TestGlobalContext.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	expectedBindings := map[string]string{
		"ESC": "closeModal",
	}

	globalContext := (*config)["Global"]
	for key, command := range expectedBindings {
		assert.Equal(t, command, globalContext.Bindings[key], "Binding for key %s should be %s", key, command)
	}
}

func TestValidContextAdd(t *testing.T) {
	configPath := "../testdata/TestValidContextAdd.toml"
	config, err := keybinding.LoadConfig(configPath)

	// No error should be present because the context_add is correct.
	assert.NoError(t, err, "Config should load without error for a valid context_add")
	assert.NotNil(t, config, "Config should not be nil")
}

func TestInvalidContextAddType(t *testing.T) {
	configPath := "../testdata/TestInvalidContextAddType.toml"
	_, err := keybinding.LoadConfig(configPath)

	// The test expects an error, as the TOML file has an invalid type for `context_add`
	assert.Error(t, err, "Config should return an error for invalid context_add type")
	assert.Contains(t, err.Error(), "incompatible types", "Error should indicate type mismatch")
}

func TestKeyBindings(t *testing.T) {
	configPath := "../testdata/TestValidConfig.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")

	defaultContext, exists := (*config)["Default"]
	assert.True(t, exists, "Default context should be present")
	assert.NotNil(t, defaultContext, "Default context should not be nil")

	assert.Equal(t, "deleteTrack", defaultContext.Bindings["d"], "Binding for key 'd' should be 'deleteTrack'")
	assert.Equal(t, "addToQueue", defaultContext.Bindings["a"], "Binding for key 'a' should be 'addToQueue'")

	queueContext, exists := (*config)["Queue"]
	assert.True(t, exists, "Queue context should be present")
	assert.NotNil(t, queueContext, "Queue context should not be nil")

	assert.Equal(t, "queue.deleteTrack", queueContext.Bindings["d"], "Binding for key 'd' in Queue should be 'queue.deleteTrack'")
	assert.Equal(t, "shuffleQueue", queueContext.Bindings["s"], "Binding for key 's' in Queue should be 'shuffleQueue'")
}

func TestBigFile(t *testing.T) {
	configPath := "../testdata/TestBigFile.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")
}

func TestContextAddErrors(t *testing.T) {
	configPath := "../testdata/TestContextAddErrors.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for context_add errors")
}

func TestContextOverrideErrors(t *testing.T) {
	configPath := "../testdata/TestContextOverrideErrors.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for context_override errors")
}

func TestEmptyFile(t *testing.T) {
	configPath := "../testdata/TestEmptyFile.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for an empty file")
}

func TestGarbageContent(t *testing.T) {
	configPath := "../testdata/TestGarbageContent.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for garbage content")
}

func TestInvalidKeyNames(t *testing.T) {
	configPath := "../testdata/TestInvalidKeyNames.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for invalid key names")
}

func TestLargeFile(t *testing.T) {
	configPath := "../testdata/TestLargeFile.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error for a large file")
	assert.NotNil(t, config, "Config should not be nil")
}

func TestMissingBindings(t *testing.T) {
	configPath := "../testdata/TestMissingBindings.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for missing bindings")
}

func TestMixedValidAndInvalid(t *testing.T) {
	configPath := "../testdata/TestMixedValidAndInvalid.toml"
	_, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error when the file contains both valid and invalid entries")
}

func TestSpecialCharacters(t *testing.T) {
	configPath := "../testdata/TestSpecialCharacters.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error for special characters")
	assert.NotNil(t, config, "Config should not be nil")
}

func TestValidConfig(t *testing.T) {
	configPath := "../testdata/TestValidConfig.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Config should load without error")
	assert.NotNil(t, config, "Config should not be nil")
	assert.NotEmpty(t, (*config)["Global"].Bindings, "Config should have contexts")
}
