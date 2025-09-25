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
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// get the shell command and trim spaces
		command = strings.TrimSpace(command)
		var cmd = strings.Fields(command)

		// Handle built-in commands
		switch cmd[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			s := strings.Join(cmd[1:], " ")
			fmt.Println(s)
		case "type":
			isBuiltin := cmd[1]
			isExe := cmd[1]

			if len(cmd) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			if _, ok := builtin[isBuiltin]; ok {
				fmt.Println(isBuiltin, "is a shell builtin")
			} else if pathExe, err := exec.LookPath(isExe); err == nil {
				fmt.Printf("%v is %v\n", isExe, pathExe)
			} else {
				fmt.Println(isBuiltin + ": not found")
			}
		default:
			fmt.Println(cmd[0] + ": command not found")
		}
	}
}
