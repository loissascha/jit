package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Username string            `json:"username"`
	Remote   Remote            `json:"remote"`
	Branches map[string]Branch `json:"branches"`
}

type Remote struct {
	Url string `json:"url"`
}

type Branch struct {
}

func LoadMainConfigFile(basepath string) *Config {
	return &Config{}
}

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
