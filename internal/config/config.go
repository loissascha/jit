package config

import "path/filepath"

func GetMainConfigFile(basepath string) string {
	return filepath.Join(basepath, ".jit/config")
}
