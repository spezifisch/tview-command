package keybinding

import "github.com/spezifisch/tview-command/types"

func ValidateConfig(config types.Config) error {
	return DetectCycleAndValidate(config)
}
