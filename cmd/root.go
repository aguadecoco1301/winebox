/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without subcommands.
var rootCmd = &cobra.Command{
	Use:   "winebox",
	Short: "Manage Wine prefixes and applications",
	Long: `Winebox is a CLI tool to create, delete, and manage Wine prefixes.

It allows you to maintain isolated environments for each Windows application
and generates .desktop files for installed applications so they can be launched
from your desktop environment.`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Disable Cobra's default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
