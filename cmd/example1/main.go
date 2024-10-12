//go:build example
// +build example

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spezifisch/tview-command/keybinding"
)

func main() {
	app := tview.NewApplication()

	// Load the config.toml file for keybindings
	configPath := "config_example1.toml"
	config, err := keybinding.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load keybinding config: %v", err)
	}

	// Create a text view for displaying the queue screen
	queueScreen := tview.NewTextView().
		SetText("Queue Screen").
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw() // Redraw the app on changes
		})

	printHelp := func() {
		// Add some instructions to the screen
		fmt.Fprintf(queueScreen, "\nHELP: Press 'd' to delete a track, 'a' to add to queue, 's' to shuffle, 'm' to move track, SPACE opens command palette.\nEXITS: Ctrl-q to quit, 'Q' to force-quit, Ctrl-C to *really* force-quit.\n")
	}
	fmt.Fprintf(queueScreen, "\nPRESS any key to start demo\n")

	// Handle global key events
	counter := 0
	queueScreen.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		counter++

		// This is our tview-command context[TM], it should always be set to the current context, e.g. by the views' "on focus" event.
		contextKey := "Queue"
		context := (*config)[contextKey]

		// Override on 'Q' to ensure quitting the application when 'Q' is pressed
		if event.Rune() == 'Q' {
			app.Stop()
			println("You quit using hard override 'Q'.")
			return nil
		}

		// Handle printable characters
		keyName := ""
		if event.Rune() != 0 {
			r := event.Rune()
			if r >= 32 && r <= 126 {
				// Printable ASCII range (basic characters)
				keyName = string(r)
			} else {
				// Non-printable characters: represent as hex
				keyName = fmt.Sprintf("0x%04X", r)
			}
		}

		if keyName == "" || strings.HasPrefix(keyName, "0x") {
			// Handle non-printable keys (e.g., ESC, Enter, Ctrl)
			switch event.Key() {
			case tcell.KeyESC:
				keyName = "ESC"
			case tcell.KeyEnter:
				keyName = "Enter"
			case tcell.KeyCtrlC:
				app.Stop() // Force quit with Ctrl-C
				println("You quit using hard override '^C'.")
				return nil
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

		// Output simple response
		queueScreen.Clear()
		fmt.Fprintf(queueScreen, `Context key: %s
Event received!
Event counter: %d

Key: %s (length: %d)
Triggered action: `, contextKey, counter, keyName, len(keyName))

		if action, ok := context.Bindings[keyName]; ok {
			fmt.Fprintf(queueScreen, "%s", action)
		} else {
			fmt.Fprintf(queueScreen, "(shortcut not bound)")
		}

		println()
		printHelp()

		return event
	})

	// Start the application with the queue screen as the root
	if err := app.SetRoot(queueScreen, true).Run(); err != nil {
		panic(err)
	}
}
