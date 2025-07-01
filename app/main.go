package main

import (
	"fmt"
	"os"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
	"cd":   true,
}

func main() {
	for {

		currentPath, err := os.Getwd()
		errorCheck(err)

		fmt.Fprint(os.Stdout, currentPath+"$ ")

		input, err := readStdio()
		errorCheck(err)

		cmd, args := input[0], input[1:]
		parseCommand(cmd, args)
	}
}
