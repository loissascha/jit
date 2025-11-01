package commands

import (
	"fmt"
	"jit/internal/config"
)

func InitCommand() {
	if config.IsJitProject(".") {
		fmt.Println("There is already a jit project initialized. Try jit help.")
		return
	}

	var username string

	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Println("")

	basicConfig := config.NewConfig()
	basicConfig.Username = username

	err := config.WriteConfigFile(".", basicConfig)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating config. Please check file permissions!")
		return
	}
}
