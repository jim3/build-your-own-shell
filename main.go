package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtin = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"cd":   true,
	"pwd":  true,
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Trim spaces and get shell command
		command = strings.TrimSpace(command)
		var cmd = strings.Fields(command)

		// Handle builtin commands and executables
		switch cmd[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			s := strings.Join(cmd[1:], " ")
			fmt.Println(s)
		case "type":
			isBuiltin := cmd[1]

			if len(cmd) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			if _, ok := builtin[isBuiltin]; ok {
				fmt.Println(isBuiltin, "is a shell builtin")
			} else if exe, err := exec.LookPath(isBuiltin); err == nil {
				fmt.Printf("%v is %v\n", isBuiltin, exe)
			} else {
				fmt.Println(isBuiltin + ": not found")
			}
		default:
			// Run external programs with arguments
			if _, err := exec.LookPath(cmd[0]); err == nil {
				command := exec.Command(cmd[0], cmd[1:]...)
				command.Stdout = os.Stdout
				command.Stderr = os.Stderr

				if err := command.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
			} else {
				fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd[0])
			}
		}
	}
}
