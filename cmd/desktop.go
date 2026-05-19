/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aguadecoco1301/winebox/internal/appstore"
	"github.com/aguadecoco1301/winebox/internal/desktop"
	"github.com/aguadecoco1301/winebox/internal/prefix"
	"github.com/spf13/cobra"
)

// desktopCmd represents the desktop command
var desktopCmd = &cobra.Command{
	Use:   "desktop",
	Short: "Manage desktop integration",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var desktopGenerateCmd = &cobra.Command{
	Use:   "generate <prefix> [app]",
	Short: "Generate a desktop entry",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]

		var appName string

		if len(args) == 2 {
			appName = args[1]

			if _, err := appstore.Load(appName); err != nil {
				fmt.Println("ERROR: app does not exists")
				return
			}
		}

		if _, err := os.Stat(prefix.Path(prefixName)); err != nil {
			fmt.Println("ERROR: prefix does not exists")
			return
		}

		execCommand := "winebox run " + prefixName
		if appName != "" {
			execCommand += " " + appName
		}
		displayName := prefixName
		if appName != "" {
			displayName = appName
		}

		entry := desktop.Entry{
			Name: displayName,
			Exec: execCommand,
		}

		content := desktop.Build(entry)

		fileName := desktop.FileName(prefixName, appName)

		err := desktop.Write(fileName, content)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println(content)
	},
}

func init() {
	rootCmd.AddCommand(desktopCmd)
	desktopCmd.AddCommand(desktopGenerateCmd)
}
