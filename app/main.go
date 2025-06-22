package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EvaluateInput(command string) (string, string, bool) {
	command = strings.TrimSpace(command)
	return strings.Cut(command, " ")
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		command, args, found := EvaluateInput(command)

		if found {
			if command == "echo" {
				fmt.Println(args)
			} else {
				fmt.Println(command + ": command not found")
			}
		} else {

			if command == "exit" {
				os.Exit(0)
			} else {
				fmt.Println(command + ": command not found")
			}
		}
	}
}
