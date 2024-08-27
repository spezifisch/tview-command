package keybinding

import (
	"fmt"
	"testing"

	"github.com/expr-lang/expr"
)

func setupEnvironment() map[string]interface{} {
	// Mock functions simulating music player commands
	return map[string]interface{}{
		"AddTrackToPlaylist": func(trackID string) string {
			return fmt.Sprintf("Track %s added to playlist", trackID)
		},
		"favoriteTrack": func(trackID string) string {
			return fmt.Sprintf("Track %s favorited", trackID)
		},
		"CurrentTrackID": "12345", // Simulate the current track ID
	}
}

func executeExpression(key string, exprString string, env map[string]interface{}) (string, error) {
	program, err := expr.Compile(exprString, expr.Env(env))
	if err != nil {
		return "", fmt.Errorf("Failed to compile script for key %s: %v", key, err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		return "", fmt.Errorf("Error executing script for key %s: %v", key, err)
	}
	return fmt.Sprintf("%v", output), nil
}
func TestExpressionExecution(t *testing.T) {
	env := setupEnvironment()
	tests := []struct {
		key    string
		expr   string
		output string
	}{
		{"A", "AddTrackToPlaylist(CurrentTrackID)", "Track 12345 added to playlist"},
		{"Y", "favoriteTrack(CurrentTrackID)", "Track 12345 favorited"},
	}

	for _, tt := range tests {
		output, err := executeExpression(tt.key, tt.expr, env)
		if err != nil {
			t.Errorf("Error in test: %v", err)
		}
		if output != tt.output {
			t.Errorf("Expected %q, got %q", tt.output, output)
		}
	}
}

func TestKeybindingIntegrationWithConfig(t *testing.T) {
	// Load the keybinding configuration
	config, err := LoadConfig("../testdata/TestValidDictBindingsSyntax.toml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	env := setupEnvironment()

	expectedOutputs := map[string]string{
		"A": "Track 12345 added to playlist",
		"B": "Track 12345 favorited",
	}

	if context, exists := (*config)["Default"]; exists {
		for key, command := range context.Bindings {
			output, err := executeExpression(key, command, env)
			if err != nil {
				t.Errorf("Error executing command %s: %v", command, err)
			} else if output != expectedOutputs[key] {
				t.Errorf("Output for key %s: expected %q, got %q", key, expectedOutputs[key], output)
			}
		}
	} else {
		t.Fatalf("Default context not found in config")
	}
}
