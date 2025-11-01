package branches

import (
	"encoding/json"
	"fmt"
	"jit/internal/config"
	"os"
	"path/filepath"
)

type Branch struct {
	CurrentCommit string `json:"current_commit"`
}

// filepath: ./.jit/branches/[branchname] -> holds the branch data

func GetBranch(basepath string, name string) (*Branch, error) {
	path := filepath.Join(config.GetConfigDirPath(basepath), "branches/"+name)
	fmt.Println("path:", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var branch Branch

	err = json.Unmarshal(content, &branch)
	if err != nil {
		return nil, err
	}

	return &branch, nil
}
