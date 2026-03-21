package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// PrintVersion outputs the flow plugin version.
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "wave-flow %s\n", GetVersion())
}

// GetVersion returns the flow plugin version by reading the Waveplugin file.
func GetVersion() string {
	// Try to find the binary path to locate Waveplugin
	exePath, err := os.Executable()
	if err != nil {
		return "unknown"
	}

	// Waveplugin is in the same directory as the binary
	wavepluginPath := filepath.Join(filepath.Dir(exePath), "Waveplugin")
	if _, err := os.Stat(wavepluginPath); err == nil {
		if content, err := os.ReadFile(wavepluginPath); err == nil {
			// Parse version from Waveplugin
			for _, line := range strings.Split(string(content), "\n") {
				if strings.HasPrefix(strings.TrimSpace(line), "version") {
					parts := strings.Split(line, "=")
					if len(parts) == 2 {
						return strings.Trim(strings.TrimSpace(parts[1]), `"`)
					}
				}
			}
		}
	}

	return "unknown"
}
