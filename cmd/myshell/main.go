package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
}
