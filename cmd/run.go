/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aguadecoco1301/winebox/internal/appstore"
	"github.com/aguadecoco1301/winebox/internal/prefix"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <prefix> [app]",
	Short: "Run an application inside a Wine prefix",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]

		var appName string

		if len(args) == 2 {
			appName = args[1]
		}

		prefixPath := prefix.Path(prefixName)

		if _, err := os.Stat(prefixPath); err != nil {
			fmt.Println("ERROR: prefix does not exists")
		}

		store, err := appstore.Load(prefixPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		app, err := store.Resolve(appName)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if !filepath.IsAbs(app.Path) {
			app.Path = filepath.Join(prefixPath, app.Path)
		}

		if _, err := os.Stat(app.Path); err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		wineCmd := exec.Command("wine", app.Path)
		wineCmd.Env = append(os.Environ(), "WINEPREFIX="+prefixPath)

		wineCmd.Stdin = os.Stdin
		wineCmd.Stdout = os.Stdout
		wineCmd.Stderr = os.Stderr

		fmt.Printf("Executing:\n> %s\n\n", wineCmd)

		if err := wineCmd.Run(); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
