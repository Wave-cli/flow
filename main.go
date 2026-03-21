// wave-flow is a wave CLI plugin for development workflow automation.
//
// It reads command definitions from the [flow] section of a Wavefile
// and executes them with optional environment variables, success/failure
// callbacks, and watch mode.
//
// Usage:
//
//	wave flow <command> [args...]
//	wave flow --list
//	wave flow --version
package main

import (
	"os"

	"github.com/wave-cli/wave-flow/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
