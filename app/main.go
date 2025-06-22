package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TrimCommand(command string) string {
	return strings.TrimSpace(command)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		command = TrimCommand(command)

		if command == "exit 0" {
			os.Exit(0)
		}
		fmt.Println(command + ": command not found")
	}
}
