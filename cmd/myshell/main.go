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

		commands := []string{"echo", "exit", "type", "pwd", "cd"}

		switch cmdName {
		case "exit":
			os.Exit(0)

		case "echo":
			fmt.Println(strings.Join(cmdArgs, " "))

		case "type":
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

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("pwd: cannot get current directory")
				continue
			}
			fmt.Println(dir)

		case "cd":
			if len(cmdArgs) == 0 {
				fmt.Println("cd: missing operand")
				continue
			}
			newDir := cmdArgs[0]
			if strings.HasPrefix(newDir, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println("cd: cannot determine home directory")
					continue
				}
				if newDir == "~" {
					newDir = homeDir
				} else {
					newDir = filepath.Join(homeDir, newDir[1:])
				}
			}
			err := os.Chdir(newDir)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", newDir)
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
