package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

func evaluateInput(command string) []string {
	return strings.Fields(command)
}

func readStdio() ([]string, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return []string{}, fmt.Errorf("input error: %v", err)
	}
	return evaluateInput(input), nil
}

func handleExit() {
	os.Exit(0)
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleType(cmd string) {
	if v := builtins[cmd]; v {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else if file, exist := findBinFile(cmd); exist {
		fmt.Printf("%s is %s\n", cmd, file)
	} else {
		fmt.Printf("%s not found\n", cmd)
	}
}

func findBinFile(bin string) (string, bool) {
	paths := os.Getenv("PATH")
	for _, path := range strings.Split(paths, ":") {
		file := path + "/" + bin
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}
	return "", false
}

// TODO pass arg checks to respective handler functions

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := readStdio()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		cmd := input[0]
		args := input[1:]

		switch cmd {
		case "exit":
			if len(args) != 0 {
				fmt.Println("Too many arguments for exit: " + strings.Join(args, " "))
			} else {
				handleExit()
			}

		case "echo":
			if len(args) == 0 {
				fmt.Println("Not enough arguments for echo")
			} else {
				handleEcho(args)
			}

		case "type":
			if len(args) > 1 {
				fmt.Println("Too many arguments for type:  " + strings.Join(args, " "))
			} else if len(args) < 1 {
				fmt.Println("Not enough arguments for type")
			} else {
				handleType(args[0])
			}
		}
	}
}
