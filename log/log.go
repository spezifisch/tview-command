package log

import "fmt"

var logPrefix = "tview-command: "
var logHandler func(string)

// SetLogHandler allows `stmps` to provide a custom logging function.
func SetLogHandler(handler func(string)) {
	logHandler = handler
}

// SetLogPrefix set the prefix to prepend directly to logged lines.
func SetLogPrefix(prefix string) {
	logPrefix = prefix
}

// Use LogMessage wherever logging is needed
func LogMessage(msg string) {
	if logHandler != nil {
		logHandler(logPrefix + msg)
	} else {
		fmt.Println(msg) // Default to standard output if no handler is set
	}
}
