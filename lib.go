package tviewcommand

import (
	"github.com/spezifisch/tview-command/context"
	"github.com/spezifisch/tview-command/keybinding"
)

// Re-export functions, types, and variables from keybinding package
// so it's easier to use for other packages.
var (
	LoadConfig     = keybinding.LoadConfig
	ValidateConfig = keybinding.ValidateConfig
)

type (
	Config  = keybinding.Config
	Context = context.Context
)
