package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	tviewcommand "github.com/spezifisch/tview-command"
	"github.com/spezifisch/tview-command/keybinding"
	tcLog "github.com/spezifisch/tview-command/log"
)

func main() {
	app := tview.NewApplication()

	logger, loggerCleanup := setupLogger()
	defer loggerCleanup()

	// Set the library's log handler to use logrus for output
	tcLog.SetLogHandler(func(msg string) {
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
	queueScreen.
		SetBorder(true).
		SetTitle("Queue Context Widget")

	// Create the text view for displaying the config dump (resolved inheritance)
	configDumpScreen := tview.NewTextView()
	configDumpScreen.
		SetText("").
		SetDynamicColors(true).
		SetBorder(true).
		SetTitle("Config Dump (Resolved Inheritance)")

	// Function to update the config dump screen dynamically
	updateConfigDump := func() {
		// Pretty-print the config struct
		configDumpBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Errorf("Failed to pretty-print config: %v", err)
			configDumpScreen.SetText(fmt.Sprintf("Error: %v", err))
		} else {
			configDumpScreen.SetText(string(configDumpBytes))
		}
	}

	// Function to show example's help
	printHelp := func() {
		// Add some instructions to the screen
		fmt.Fprintf(queueScreen, "\n\nHELP: Press 'd' to delete a track, 'a' to add to queue, 's' to shuffle, 'm' to move track, SPACE opens command palette.\n\nEXITS: Ctrl-q to quit, 'Q' to force-quit, Ctrl-C to *really* force-quit.\n")
	}
	fmt.Fprintf(queueScreen, "\nPRESS any key to start demo\n")

	// Handle global key events for queueScreen and update config dump
	counter := 0
	queueScreen.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		counter++
		logger.Debugf("Key event received: counter=%d", counter)

		// Resolve the keybindings for the Queue context dynamically
		contextKey := "Queue"
		currentContext := (*config)[contextKey]
		logger.Debugf("Using context: %s", contextKey)

		// Override on 'Q' to ensure quitting the application
		if event.Rune() == 'Q' {
			logger.Info("User pressed 'Q' to quit")
			app.Stop()
			logger.Info("Application quit using hard override 'Q'.")
			return nil
		}

		// Handle key input
		tcEvent := tviewcommand.FromEventKey(event)
		logger.Debugf("Got Event: %s", tview.Escape(tcEvent.String()))

		// Update screen with event details
		queueScreen.Clear()
		fmt.Fprintf(queueScreen, `Context key: %s
Event received!
Event counter: %d

Event: %s
Triggered action: `, contextKey, counter, tview.Escape(tcEvent.String()))

		// Trigger appropriate action based on the resolved keybindings
		if action, ok := currentContext.Bindings[tcEvent.KeyName]; ok {
			fmt.Fprintf(queueScreen, "%s", tview.Escape(action))
			logger.Infof("Triggered action for key '%s': %s", tcEvent.KeyName, action)
		} else {
			fmt.Fprintf(queueScreen, "(shortcut not bound)")
			logger.Warnf("No action bound for key: %s", tcEvent.KeyName)
		}
		printHelp()

		// Update the config dump screen with the latest config state
		updateConfigDump()

		return event
	})

	// Split the screen into two main sections: left (queueScreen) and right (file and config dump)
	mainSplit := tview.NewFlex().
		AddItem(queueScreen, 0, 1, true).     // Left side (queueScreen)
		AddItem(configDumpScreen, 0, 1, true) // Right side (vsplit)

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
