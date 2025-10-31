package commands

import (
	"fmt"
	"jit/internal/status"
)

func StatusCommand() {
	if !status.IsJitProject(".") {
		fmt.Println("Not a jit project. Use jit init.")
		return
	}
	fmt.Println("Status command")
}
