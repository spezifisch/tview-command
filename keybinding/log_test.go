package keybinding

import (
	"strings"
	"testing"
)

func TestLogHandler(t *testing.T) {
	var loggedMessages []string

	// Custom log handler to collect log messages
	SetLogHandler(func(msg string) {
		loggedMessages = append(loggedMessages, msg)
	})

	// Trigger logging in the keybinding package
	logMessage("Test log message")

	// Check if the log message was captured correctly
	expectedMessage := "tview-command: Test log message"
	if len(loggedMessages) != 1 || loggedMessages[0] != expectedMessage {
		t.Errorf("Expected %q, got %q", expectedMessage, strings.Join(loggedMessages, ", "))
	}

	// Reset the log handler to the default
	SetLogHandler(nil)

	// Trigger another log, which should not be captured by our handler
	logMessage("This should go to stdout")

	// Ensure that no new messages were captured
	if len(loggedMessages) != 1 {
		t.Errorf("Expected no additional messages, but got %d", len(loggedMessages))
	}
}
