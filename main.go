package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

func main() {
	if len(os.Args) < 2 {
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
		fmt.Fprintln(os.Stderr, "Unknown Command", command)
		os.Exit(3)
	}
}

func runCommand(args []string) {

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "dmux: usage: dmux run [command] [args...]")
		os.Exit(3)
	}

	idx := slices.Index(args, "--")
	if idx != -1 {
		args = slices.Delete(args, idx, idx+1)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Fprint(os.Stderr, "dmux:", err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		os.Exit(cmd.ProcessState.ExitCode())
	}

	os.Exit(cmd.ProcessState.ExitCode())

}
