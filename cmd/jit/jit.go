package main

import (
	"flag"
	"jit/internal/commands"
)

func main() {
	// status := flag.Bool("status", false, "show the current status")
	flag.Parse()
	args := flag.Args()

	for _, a := range args {
		if a == "status" {
			commands.StatusCommand()
		}
	}
}
