/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/

// winebox app list
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/aguadecoco1301/winebox/internal/appstore"
	"github.com/aguadecoco1301/winebox/internal/prefix"
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

var appMainCmd = &cobra.Command{
	Use:   "main <prefix> <path>",
	Short: "Set the main application for a prefix",
	Long:  `Defines which registered application will be used as default when running 'winebox run'.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]
		appPath := args[1]

		absPath, err := filepath.Abs(appPath)
		if err != nil {
			return
		}
		absPath = filepath.Clean(absPath)

		prefixPath := prefix.Path(prefixName)

		err = appstore.SetMain(prefixPath, absPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Main set for prefix:", prefixName)
	},
}

var appAddCmd = &cobra.Command{
	Use:   "add <prefix> <name> <path>",
	Short: "Add a new application to a prefix",
	Long: `Registers a new executable inside a Wine prefix.

The <name> is an alias used later with 'winebox run'.
The <path> is the executable path inside the prefix or relative to it.`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]
		appName := args[1]
		appPath := args[2]

		absPath, err := filepath.Abs(appPath)
		if err != nil {
			return
		}
		absPath = filepath.Clean(absPath)

		prefixPath := prefix.Path(prefixName)

		err = appstore.AddApp(prefixPath, appName, absPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("App added:", appName)
	},
}

var appEditCmd = &cobra.Command{
	Use:   "edit <prefix> <old-name> <new-name> <new-path>",
	Short: "Edit an existing application entry",
	Long: `Modifies an existing application inside a Wine prefix.

You can change both the alias name and the executable path.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]
		oldName := args[1]
		newName := args[2]
		newPath := args[3]

		prefixPath := prefix.Path(prefixName)

		absPath, err := filepath.Abs(newPath)
		if err != nil {
			return
		}
		absPath = filepath.Clean(absPath)

		err = appstore.EditApp(prefixPath, oldName, newName, absPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("App updated:", oldName)
	},
}

var appDeleteCmd = &cobra.Command{
	Use:   "delete <prefix> <name>",
	Short: "Delete an application from a prefix",
	Long: `Removes a registered application from a Wine prefix.

This only deletes the application entry from the registry.
It does NOT delete the prefix itself.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]
		appName := args[1]

		prefixPath := prefix.Path(prefixName)

		err := appstore.DeleteApp(prefixPath, appName)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("App deleted:", appName)
	},
}

func init() {
	rootCmd.AddCommand(appCmd)

	appCmd.AddCommand(appMainCmd)
	appCmd.AddCommand(appAddCmd)
	appCmd.AddCommand(appEditCmd)
	appCmd.AddCommand(appDeleteCmd)
}
