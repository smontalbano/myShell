package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
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

func handleExit(args []string) {
	if len(args) > 1 {
		fmt.Println("Too many arguments for exit: " + strings.Join(args, " "))
	} else if len(args) == 0 {
		os.Exit(0)
	} else {
		if num, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("Invalid exit code: " + args[0] + "\nExit code must be of type integer")
			return
		} else {
			os.Exit(num)
		}
	}
}

func handleEcho(args []string) {
	if len(args) == 0 {
		fmt.Println("Not enough arguments for echo")
	} else {
		fmt.Println(strings.Join(args, " "))
	}
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

func parseCommand(cmd string, args []string) {
	switch cmd {
	case "exit":
		handleExit(args)

	case "echo":
		handleEcho(args)

	case "type":
		if len(args) != 1 {
			fmt.Printf("Incorrect number of arguments for type\nExpected: 1 Received: %v\n", len(args))
		} else {
			handleType(args[0])
		}
	default:
		checkForCommand(cmd, args, os.Stdout)
	}
}

func checkForCommand(cmd string, args []string, out io.Writer) {
	if file, exists := findBinFile(cmd); exists {
		runner := exec.Command(file, args...)
		result, err := runner.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		out.Write(result)
	}
}
func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := readStdio()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		cmd, args := input[0], input[1:]
		parseCommand(cmd, args)
	}
}
