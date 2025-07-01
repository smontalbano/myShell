package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func handlePwd() {
	currentDirectory, err := os.Getwd()
	errorCheck(err)
	fmt.Println(currentDirectory)
}

func handleCd(path []string) {
	home, err := os.UserHomeDir()
	errorCheck(err)

	if len(path) > 1 {
		fmt.Printf("Incorrect number of arguments for cd\nExpected: 1 Received: %v\n", len(path))
	} else if len(path) == 0 {
		err = os.Chdir(home)
		errorCheck(err)
	} else if string(path[0][0]) == "~" {
		err = os.Chdir(home + path[0][1:])
		errorCheck(err)
	} else {
		err := os.Chdir(path[0])
		errorCheck(err)
	}
}
