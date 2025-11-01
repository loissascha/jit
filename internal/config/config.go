package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Username string            `json:"username"`
	Remotes  map[string]Remote `json:"remotes"`
	Branches map[string]Branch `json:"branches"`
}

type Remote struct {
	Url string `json:"url"`
}

type Branch struct {
	Active bool `json:"active"`
}

func NewConfig() *Config {
	return &Config{
		Username: "",
		Remotes:  map[string]Remote{},
		Branches: map[string]Branch{"main": {Active: true}},
	}
}

func LoadMainConfigFile(basepath string) (*Config, error) {
	path := GetMainConfigFilePath(basepath)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c, err := mapConfig(content)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func WriteConfigFile(basepath string, config *Config) error {
	path := GetMainConfigFilePath(basepath)

	jsonData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func mapConfig(content []byte) (*Config, error) {
	var res Config
	err := json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
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
