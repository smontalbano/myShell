package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
	"cd":   true,
}

func evaluateInput(command string) []string {
	return strings.Fields(command)
}

func readStdio() ([]string, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	errorCheck(err)
	return evaluateInput(input), nil
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
	case "pwd":
		handlePwd()

	case "cd":
		handleCd(args)

	default:
		checkForCommand(cmd, args, os.Stdout)
	}
}

func checkForCommand(cmd string, args []string, out io.Writer) {
	if file, exists := findBinFile(cmd); exists {
		runner := exec.Command(file, args...)
		result, err := runner.CombinedOutput()
		errorCheck(err)
		out.Write(result)
	}
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
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
