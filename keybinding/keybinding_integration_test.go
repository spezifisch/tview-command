package keybinding

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

func TestIntegrationWithExpr(t *testing.T) {
	env := map[string]interface{}{
		"AddTrackToPlaylist": func(trackID string) bool {
			fmt.Printf("Track %s added to playlist\n", trackID)
			return true
		},
		"ToggleFavoriteTrack": func(trackID string) bool {
			fmt.Printf("Track %s favorite status toggled\n", trackID)
			return true
		},
		"CurrentTrackID": "1234", // Example Track ID
	}

	configCommands := map[string]string{
		"A": "AddTrackToPlaylist(CurrentTrackID)",
		"y": "ToggleFavoriteTrack(CurrentTrackID)",
	}

	compiledPrograms := make(map[string]*vm.Program)
	for key, script := range configCommands {
		program, err := expr.Compile(script, expr.Env(env))
		if err != nil {
			t.Fatalf("Failed to compile script for key %s: %v", key, err)
		}
		compiledPrograms[key] = program
	}

	tests := []struct {
		key     string
		trackID string
		output  string
	}{
		{"A", "track-id-123", "Track track-id-123 added to playlist\n"},
		{"y", "track-id-456", "Track track-id-456 favorite status toggled\n"},
	}

	for _, tt := range tests {
		env["CurrentTrackID"] = tt.trackID
		if program, exists := compiledPrograms[tt.key]; exists {
			output := captureOutput(func() {
				_, err := expr.Run(program, env)
				if err != nil {
					t.Errorf("Error executing script for key %s: %v", tt.key, err)
				}
			})
			if output != tt.output {
				t.Errorf("Expected %q, got %q", tt.output, output)
			}
		}
	}
}

// captureOutput is a utility function to capture printed output during tests.
func captureOutput(f func()) string {
	var buf bytes.Buffer
	writer := io.MultiWriter(os.Stdout, &buf)
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	io.Copy(writer, r)
	return buf.String()
}
