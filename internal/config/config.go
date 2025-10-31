package config

import (
	"os"
	"path/filepath"
)

func GetMainConfigFilePath(basepath string) string {
	return filepath.Join(basepath, ".jit/config")
}

func IsJitProject(path string) bool {
	configPath := GetMainConfigFilePath(path)
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
