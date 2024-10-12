package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spezifisch/tview-command/keybinding"
)

func main() {
	app := tview.NewApplication()

	// Configure logrus for colored output
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set the library's log handler to use logrus for output
	keybinding.SetLogHandler(func(msg string) {
		logger.Info(msg) // Sends the log message to logrus
	})

	logger.Info("Starting tview-command example1...")

	// Load the config.toml file for keybindings
	configPath := "config_example1.toml"
	logger.Infof("Loading keybinding config from %s", configPath)
	config, err := keybinding.LoadConfig(configPath)
	if err != nil {
		logger.Fatalf("Failed to load keybinding config: %v", err)
	}
	logger.Info("Keybinding config loaded successfully")

	// Create a text view for displaying the queue screen
	queueScreen := tview.NewTextView().
		SetText("Queue Screen").
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw() // Redraw the app on changes
		})

	logger.Debug("Queue screen created")

	// Helper to print help instructions
	printHelp := func() {
		fmt.Fprintf(queueScreen, "\n\nHELP: Press 'd' to delete a track, 'a' to add to queue, 's' to shuffle, 'm' to move track, SPACE opens command palette.\nEXITS: Ctrl-q to quit, 'Q' to force-quit, Ctrl-C to *really* force-quit.\n")
	}
	queueScreen.SetText("\nPRESS any key to start demo\n")

	// Handle global key events
	counter := 0
	queueScreen.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		counter++
		logger.Debugf("Key event received: counter=%d", counter)

		// Use the current context
		contextKey := "Queue"
		context := (*config)[contextKey]
		logger.Debugf("Using context: %s", contextKey)

		// Override on 'Q' to ensure quitting the application
		if event.Rune() == 'Q' {
			logger.Info("User pressed 'Q' to quit")
			app.Stop()
			logger.Info("Application quit using hard override 'Q'.")
			return nil
		}

		// Handle key input
		keyName := getKeyName(event)

		logger.Debugf("Key pressed: %s", keyName)

		// Update screen with event details
		queueScreen.Clear()
		fmt.Fprintf(queueScreen, `Context key: %s
Event received!
Event counter: %d

Key: %s (length: %d)
Triggered action: `, contextKey, counter, keyName, len(keyName))

		// Trigger appropriate action
		if action, ok := context.Bindings[keyName]; ok {
			fmt.Fprintf(queueScreen, "%s", action)
			logger.Infof("Triggered action for key '%s': %s", keyName, action)
		} else {
			fmt.Fprintf(queueScreen, "(shortcut not bound)")
			logger.Warnf("No action bound for key: %s", keyName)
		}

		printHelp()

		return event
	})

	// Start the application with the queue screen as the root
	logger.Info("Starting the TUI application...")
	if err := app.SetRoot(queueScreen, true).Run(); err != nil {
		logger.Fatalf("Application crashed: %v", err)
	}
	logger.Info("Application stopped")
}

// getKeyName handles printable and non-printable key inputs
func getKeyName(event *tcell.EventKey) string {
	keyName := ""
	if event.Rune() != 0 {
		r := event.Rune()
		if r >= 33 && r <= 126 {
			// Printable ASCII range (basic characters)
			keyName = string(r)
		} else if r == 32 {
			keyName = "SPC"
		} else {
			// Non-printable characters: represent as hex
			keyName = fmt.Sprintf("0x%04X", r)
		}
	} else {
		// Handle non-printable keys (e.g., ESC, Enter, Ctrl)
		switch event.Key() {
		case tcell.KeyESC:
			keyName = "ESC"
		case tcell.KeyEnter:
			keyName = "Enter"
		case tcell.KeyCtrlC:
			keyName = "CTRL-C"
		case tcell.KeyCtrlV:
			keyName = "CTRL-V"
		case tcell.KeyCtrlX:
			keyName = "CTRL-X"
		case tcell.KeyCtrlZ:
			keyName = "CTRL-Z"
		case tcell.KeyCtrlQ:
			keyName = "CTRL-Q"
		default:
			keyName = fmt.Sprintf("keycode-%d", event.Key())
		}
	}
	return keyName
}
