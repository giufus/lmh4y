package main

import (
	"github.com/giufus/lmh4y/cmd"
)

func main() {
	// Execute the root command from the cmd package.
	// Cobra handles all the command-line parsing from here.
	cmd.Execute()
}
