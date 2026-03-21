package cmd

import (
	"fmt"
	"io"

	"github.com/wave-cli/wave-flow/internal/flow"
)

// ListCommands outputs all available flow commands.
// Returns 0 on success.
func ListCommands(config map[string]any, w io.Writer) int {
	cmds := flow.ListCommands(config)
	if len(cmds) == 0 {
		fmt.Fprintln(w, "No flow commands defined. Add commands to the [flow] section of your Wavefile.")
		return 0
	}
	fmt.Fprintln(w, "Available flow commands:")
	for _, name := range cmds {
		fmt.Fprintf(w, "  %s\n", name)
	}
	return 0
}
