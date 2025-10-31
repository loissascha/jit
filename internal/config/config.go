package config

import "path/filepath"

func GetMainConfigFilePath(basepath string) string {
	return filepath.Join(basepath, ".jit/config")
}
