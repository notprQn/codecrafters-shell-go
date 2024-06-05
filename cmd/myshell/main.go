package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		commands := []string{"echo", "exit"}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cmd = strings.TrimSpace(cmd)

		switch {
		case cmd == "exit 0":
			os.Exit(0)

		case strings.HasPrefix(cmd, "echo "):
			output := strings.TrimPrefix(cmd, "echo ")
			fmt.Printf("%s\n", output)

		case strings.HasPrefix(cmd, "type "):
			output := strings.TrimPrefix(cmd, "type ")

			if slices.Contains(commands, output) {
				fmt.Printf("%s is a shell builtin\n", output)
			} else {
				fmt.Printf("%s not found\n", output)
			}

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
