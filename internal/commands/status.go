package commands

import (
	"fmt"
	"jit/internal/branches"
	"jit/internal/config"
	"os"
)

func StatusCommand() {
	if !config.IsJitProject(".") {
		fmt.Println("Not a jit project. Use jit init.")
		return
	}

	c, err := config.LoadMainConfigFile(".")
	if err != nil {
		fmt.Println("Error reading config file. Check file permissions.")
		return
	}

	activeBranch := ""
	for k, v := range c.Branches {
		if v.Active {
			activeBranch = k
		}
	}

	if activeBranch == "" {
		fmt.Println("No active branch...")
		return
	}

	fmt.Println("Active Branch:", activeBranch)

	b, err := branches.GetBranch(".", activeBranch)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("all files are new as there is no file yet!")
			return
		}
		fmt.Println("Error reading branch file.", err)
		return
	}
	fmt.Println(b)
}
