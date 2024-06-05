package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

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
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
