package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2{
		fmt.Fprintln(os.Stderr, "dmux: usage: dmux [args]")
		os.Exit(3)
	}



	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println("0.1.0")
	case "help":
		fmt.Println("help")
	case "run":
		shell(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "Unknown Command $s\n", command)
		os.Exit(3)
	}
}

func shell(args []string){

	output, err := exec.Command("sh", "-c", strings.Join(args[1:], " ")).Output()

	if err != nil {
		fmt.Fprintln(os.Stderr, "$s\n", err)
		os.Exit(3)
	}

	fmt.Printf("%s\n", output)
}
