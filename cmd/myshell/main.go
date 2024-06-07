package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmdLine, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cmdLine = strings.TrimSpace(cmdLine)
		cmdFields := strings.Fields(cmdLine)

		if len(cmdFields) == 0 {
			continue
		}

		cmdName := cmdFields[0]
		cmdArgs := cmdFields[1:]

		commands := []string{"echo", "exit", "type"}

		switch {
		case cmdName == "exit":
			os.Exit(0)

		case cmdName == "echo":
			fmt.Println(strings.Join(cmdArgs, " "))

		case cmdName == "type":
			if len(cmdArgs) == 0 {
				fmt.Println("type: missing operand")
				continue
			}

			output := cmdArgs[0]
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
			path := os.Getenv("PATH")
			paths := strings.Split(path, ":")
			found := false
			var fullPath string

			for _, p := range paths {
				fullPath = filepath.Join(p, cmdName)
				if _, err := os.Stat(fullPath); err == nil {
					found = true
					break
				}
			}

			if found {
				execCmd := exec.Command(fullPath, cmdArgs...)
				execCmd.Stdout = os.Stdout
				execCmd.Stderr = os.Stderr
				err = execCmd.Run()
				if err != nil {
					fmt.Printf("%s: failed to execute\n", cmdName)
				}
			} else {
				fmt.Printf("%s: not found\n", cmdName)
			}
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
