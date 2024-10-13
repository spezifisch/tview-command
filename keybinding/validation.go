package keybinding

import "github.com/spezifisch/tview-command/types"

func ValidateConfig(config types.Config) error {
	// space for other checks in the future
	return DetectCycleAndValidate(config)
}
