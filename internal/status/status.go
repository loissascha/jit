package status

import (
	"jit/internal/config"
	"os"
)

func IsJitProject(path string) bool {
	configPath := config.GetMainConfigFilePath(path)
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
