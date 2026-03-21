// Package cmd implements the wave-flow CLI commands.
package cmd

import (
	"fmt"
	"io"

	"github.com/wave-cli/wave-core/pkg/sdk"
	"github.com/wave-cli/wave-flow/internal/flow"
)

// formatError writes a plain text error to stderr.
// Format: "code: message\ndetails" (if details present)
func formatError(w io.Writer, code, message, details string) {
	if details != "" {
		fmt.Fprintf(w, "%s: %s\n%s\n", code, message, details)
	} else {
		fmt.Fprintf(w, "%s: %s\n", code, message)
	}
}

// Run is the main entry point for the flow plugin.
// It reads config from r, processes args, writes output to stdout/stderr,
// and returns an exit code.
func Run(args []string, r io.Reader, stdout, stderr io.Writer) int {
	// Handle -h / --help flag first (before reading config)
	if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
		PrintHelp(stdout)
		return 0
	}

	// Handle --version flag
	if len(args) > 0 && (args[0] == "--version" || args[0] == "-v") {
		PrintVersion(stdout)
		return 0
	}

	// Read config via SDK
	cfg, err := sdk.ReadConfigFrom(r)
	if err != nil {
		formatError(stderr, "flow-config-error", "failed to read config", err.Error())
		return 1
	}
	config := cfg.Raw()

	// Handle --list flag
	if len(args) > 0 && (args[0] == "--list" || args[0] == "-l") {
		return ListCommands(config, stdout)
	}

	// Require a command name
	if len(args) == 0 {
		fmt.Fprintln(stdout, "No command specified. Run 'wave flow --help' for usage.")
		return 0
	}

	cmdName := args[0]

	// Resolve and execute the command
	cmd, err := flow.ResolveCommand(config, cmdName)
	if err != nil {
		// Show error and tell user how to see available commands
		formatError(stderr, "flow-resolve-error", err.Error(), "Run 'wave flow --list' to see available commands.")
		return 1
	}

	return flow.RunCommand(cmd, stdout, stderr)
}
