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

	defaultContext := (*config)["Default"]
	assert.NotNil(t, defaultContext, "Default context should be present")

	assert.Equal(t, "deleteTrack", defaultContext.Bindings["d"], "Binding for key 'd' should be 'deleteTrack'")
	assert.Equal(t, "addToQueue", defaultContext.Bindings["a"], "Binding for key 'a' should be 'addToQueue'")

	queueContext := (*config)["Queue"]
	assert.NotNil(t, queueContext, "Queue context should be present")

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
	config, err := keybinding.LoadConfig(configPath)

	assert.Error(t, err, "Config should return an error for context_override errors")
	assert.Nil(t, config, "Config should be nil on error")
}

func TestEmptyFile(t *testing.T) {
	configPath := "../testdata/TestEmptyFile.toml"
	config, err := keybinding.LoadConfig(configPath)

	assert.NoError(t, err, "Empty config should not return an error")
	assert.NotNil(t, config, "Config should not be nil even if it's empty")
	assert.Equal(t, 0, len(*config), "Config should be empty")
}

func TestFileNotFound(t *testing.T) {
	nonExistentPath := "../testdata/NonExistentConfig.toml"

	config, err := keybinding.LoadConfig(nonExistentPath)

	// Check that an error is returned
	assert.Error(t, err, "Loading a non-existent file should return an error")
	assert.Contains(t, err.Error(), "no such file or directory", "Error should indicate the file was not found")

	// Config should be nil
	assert.Nil(t, config, "Config should be nil when file is not found")
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

func TestMissingBindings(t *testing.T) {
	path := "../testdata/TestMissingBindings.toml"

	config, err := keybinding.LoadConfig(path)

	assert.NoError(t, err, "Config with missing bindings should not return an error")
	assert.NotNil(t, config, "Config should not be nil even if bindings are missing")
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
