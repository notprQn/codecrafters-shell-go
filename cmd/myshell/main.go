package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		commands := []string{"echo", "exit", "type"}

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

			if contains(commands, output) {
				fmt.Printf("%s is a shell builtin\n", output)
			} else {
				path := os.Getenv("PATH")
				paths := strings.Split(path, ":")

				found := false
				for _, p := range paths {
					fullPath := filepath.Join(p, output)
					if _, err := os.Stat(fullPath); err == nil {
						fmt.Printf("%s is %s\n", output, fullPath)
						found = true
						break
					}
				}

				if !found {
					fmt.Printf("%s: not found\n", output)
				}
			}

		default:
			fmt.Printf("%s: not found\n", cmd)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
