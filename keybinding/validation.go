package keybinding

func ValidateConfig(config Config) error {
	return DetectCycleAndValidate(config)
}
