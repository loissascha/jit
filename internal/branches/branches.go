package branches

import (
	"fmt"
	"jit/internal/config"
	"path/filepath"
)

type Branch struct {
	CurrentCommit string `json:"current_commit"`
}

// filepath: ./.jit/branches/[branchname] -> holds the branch data

func GetBranch(basepath string, name string) {
	path := filepath.Join(config.GetConfigDirPath(basepath), "branches/"+name)
	fmt.Println("path:", path)
}
