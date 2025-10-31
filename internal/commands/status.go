package commands

import (
	"fmt"
	"jit/internal/config"
)

func StatusCommand() {
	if !config.IsJitProject(".") {
		fmt.Println("Not a jit project. Use jit init.")
		return
	}
	fmt.Println("Status command")
}
