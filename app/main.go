package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {

	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	exitCode := cmd.ProcessState.ExitCode()
	if exitCode > 0 {
		os.Exit(exitCode)
	}

	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
	fmt.Print(string(stdout.String()))
	fmt.Fprint(os.Stderr, stderr.String())

}
