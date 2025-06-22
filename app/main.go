package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewREader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		os.Exit(1)
	}
	fmt.Println(command[:len(command)-1] + ": command not found")
}
