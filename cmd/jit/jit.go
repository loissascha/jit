package main

import (
	"flag"
	"fmt"
)

func main() {
	// status := flag.Bool("status", false, "show the current status")
	flag.Parse()
	args := flag.Args()

	for _, a := range args {
		fmt.Println(a)
		if a == "status" {
			fmt.Println("Status command")
		}
	}
}
