package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtin = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

func main() {
	for {
		// print the prompt and wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n') // Wait for user input
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// get the shell command and trim spaces
		command = strings.TrimSpace(command)

		// put the command into a slice
		var cmd = strings.Fields(command)

		switch cmd[0] {
		case "exit":
			os.Exit(0)
			// echo the rest of the command
		case "echo":
			s := strings.Join(cmd[1:], " ")
			fmt.Println(s)
		case "type":
			// check if they forgot to add an argument
			if len(cmd) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			// check if the argument is a builtin command
			isBuiltin := cmd[1]
			if _, ok := builtin[isBuiltin]; ok {
				fmt.Println(isBuiltin, "is a 🐚 builtin")
			} else {
				fmt.Println(isBuiltin + ": not found")
			}
		default:
			fmt.Println(cmd[0] + ": command not found")
		}
	}
}
