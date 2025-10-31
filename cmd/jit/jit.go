package main

import (
	"flag"
	"fmt"
	"jit/internal/commands"
)

func main() {
	// status := flag.Bool("status", false, "show the current status")
	flag.Parse()
	args := flag.Args()

	for _, a := range args {
		switch a {
		case "help":
			fmt.Println("you need help or what? skill issue.")
		case "status":
			commands.StatusCommand()
		case "init":
			fmt.Println("init")
		}
	}
}
