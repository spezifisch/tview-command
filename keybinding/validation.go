package keybinding

import "github.com/spezifisch/tview-command/types"

func ValidateConfig(config types.Config) error {
	// space for other checks in the future

	// TODO check if key or key combinations of entries are valid

	return DetectCycleAndValidate(config)
}
