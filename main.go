package main

import (
	"fmt"
	"os"

	"github.com/chroju/awscccli/command"
)

const (
	version = "0.0.1"
)

func main() {
	o := os.Stdout
	e := os.Stderr

	command, err := command.NewCommand(version, o, e)
	if err != nil {
		fmt.Fprintln(e, err)
	}

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
