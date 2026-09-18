package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println("0.1.0")
	case "help":
		fmt.Println("help")
	default:
		fmt.Fprintln(os.Stderr, "Unknown Command $s\n", command)
		os.Exit(3)
	}
}
