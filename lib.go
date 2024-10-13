package tviewcommand

import (
	"github.com/spezifisch/tview-command/keybinding"
	"github.com/spezifisch/tview-command/log"
	"github.com/spezifisch/tview-command/types"
)

// Re-export functions, types, and variables from keybinding package
// so it's easier to use for other packages.
var (
	LoadConfig     = keybinding.LoadConfig
	ValidateConfig = keybinding.ValidateConfig

	SetLogHandler = log.SetLogHandler
	SetLogPrefix  = log.SetLogPrefix

	NewContextStack = types.NewContextStack
)

type (
	Config       = types.Config
	Context      = types.Context
	ContextStack = types.ContextStack
)
