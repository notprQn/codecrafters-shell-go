package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	fmt.Fprint(os.Stdout, "$ ")

	bufio.NewReader(os.Stdin).ReadString('\n')
}
