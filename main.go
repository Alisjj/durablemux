package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
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
		runCommand(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "Unknown Command $s\n", command)
		os.Exit(3)
	}
}

func runCommand(args []string){

	idx := slices.Index(args, "--")
	if idx != -1 {
		args = slices.Delete(args, idx, idx+1)
	}
	
	output, err := exec.Command(args[0], args[1:]...).Output()

	if err != nil {
		fmt.Fprint(os.Stderr, "$s", err)
		os.Exit(3)
	}

	fmt.Printf("%s", output)
}
