package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spezifisch/tview-command/keybinding"
)

func main() {
	app := tview.NewApplication()

	logger, loggerCleanup := setupLogger()
	defer loggerCleanup()

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

	// Create the text view for displaying the file content (./config_example1.toml)
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		logger.Fatalf("Failed to read config file: %v", err)
	}
	fileContentScreen := tview.NewTextView().
		SetText(string(fileContent)).
		SetDynamicColors(true)

	// Create the text view for displaying the config dump (resolved inheritance)
	configDumpScreen := tview.NewTextView().
		SetText("").
		SetDynamicColors(true)

	// Function to update the config dump screen dynamically
	updateConfigDump := func() {
		configDump := fmt.Sprintf("%+v", config) // Dump the config var
		configDumpScreen.SetText(configDump)
	}

	// Handle global key events for queueScreen and update config dump
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

		// Update the config dump screen with the latest config state
		updateConfigDump()

		return event
	})

	// Split the screen into two main sections: left (queueScreen) and right (file and config dump)
	mainSplit := tview.NewFlex().
		AddItem(queueScreen, 0, 1, true). // Left side (queueScreen)
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow). // Right side split vertically
									AddItem(fileContentScreen, 0, 1, false).             // Top (file content)
									AddItem(configDumpScreen, 0, 1, false), 0, 1, false) // Bottom (config dump)

	// Start the application with the main split screen as the root
	logger.Info("Starting the TUI application...")
	if err := app.SetRoot(mainSplit, true).Run(); err != nil {
		logger.Fatalf("Application crashed: %v", err)
	}
	logger.Info("Application stopped")
}

func setupLogger() (*logrus.Logger, func()) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Create a temporary file for logging
	tmpFile, err := os.CreateTemp("", "app_log_*.log")
	if err != nil {
		logger.Fatalf("Failed to create temp log file: %v", err)
	}

	// Set permissions to 0600 (owner can read and write)
	if err := os.Chmod(tmpFile.Name(), 0600); err != nil {
		logger.Fatalf("Failed to set file permissions: %v", err)
	}

	// Output to stdout the location of the log file
	fmt.Printf("logging into %s\n", tmpFile.Name())

	// Configure logrus to write logs into the temp file
	logger.SetOutput(tmpFile)

	// Define the cleanup function that will close the tmpFile
	loggerCleanup := func() {
		if err := tmpFile.Close(); err != nil {
			// fallback to default log package
			log.Printf("Failed to close log file: %v", err)
		}
	}

	// Return both the logger and the cleanup function
	return logger, loggerCleanup
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
