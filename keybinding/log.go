package keybinding

import "fmt"

var logHandler func(string)

// SetLogHandler allows `stmps` to provide a custom logging function.
func SetLogHandler(handler func(string)) {
	logHandler = handler
}

// Use logHandler wherever logging is needed
func logMessage(msg string) {
	if logHandler != nil {
		logHandler("tview-command: " + msg)
	} else {
		fmt.Println(msg) // Default to standard output if no handler is set
	}
}
