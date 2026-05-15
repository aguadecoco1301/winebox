/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/

// winebox app add program map PATH
// winebox app main program PATH
// winebox app edit program subprogram NEW_NAME NEW_PATH
package cmd

import (
	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app <command>",
	Short: "Manage applications inside Wine prefixes",
	Long: `Manage applications registered inside each Wine prefix.

This command allows you to define, edit and organize executable entries
inside a prefix so they can be launched easily using 'winebox run'.

Each prefix can contain multiple applications with one designated as main.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var appAddCmd = &cobra.Command{
	Use:   "add <prefix> <name> <path>",
	Short: "Add a new application to a prefix",
	Long: `Registers a new executable inside a Wine prefix.

The <name> is an alias used later with 'winebox run'.
The <path> is the executable path inside the prefix or relative to it.`,
	Args: cobra.ExactArgs(3),
}

var appEditCmd = &cobra.Command{
	Use:   "edit <prefix> <name> <new-name> <new-path>",
	Short: "Edit an existing application entry",
	Long: `Modifies an existing application inside a Wine prefix.

You can change both the alias name and the executable path.`,
	Args: cobra.ExactArgs(4),
}

var appMainCmd = &cobra.Command{
	Use:   "main <prefix> <path>",
	Short: "Set the main application for a prefix",
	Long: `Defines the default executable that will be launched when no
application name is provided in 'winebox run'.`,
	Args: cobra.ExactArgs(2),
}

var appDeleteCmd = &cobra.Command{
	Use:   "delete <prefix> <name>",
	Short: "Delete an application from a prefix",
	Long: `Removes a registered application from a Wine prefix.

This only deletes the application entry from the registry.
It does NOT delete the prefix itself.`,
	Args: cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(appCmd)

	appCmd.AddCommand(appAddCmd)
	appCmd.AddCommand(appEditCmd)
	appCmd.AddCommand(appMainCmd)
	appCmd.AddCommand(appDeleteCmd)
}
