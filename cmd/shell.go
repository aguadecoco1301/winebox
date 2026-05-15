/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell <prefix-name>",
	Args:  cobra.ExactArgs(1),
	Short: "Open an interactive shell inside a Wine prefix",
	Long: `Opens an interactive shell inside a selected Wine prefix environment.

This command sets the WINEPREFIX environment variable and starts a new shell session,
allowing you to run installers and applications directly inside the prefix.

It is intended for installing software and manually launching programs during setup.
For daily usage of installed applications, prefer 'winebox run'.`,
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]

		prefixPath := filepath.Join(os.Getenv("HOME"), ".winebox", "prefixes", prefixName, "prefix")

		if _, err := os.Stat(prefixPath); err != nil {
			fmt.Println("ERROR: prefix does not exist:", err)
			return
		}

		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/bash"
		}

		historyFile := filepath.Join(os.Getenv("HOME"), ".winebox", "prefixes", prefixName, "history")

		cmdShell := exec.Command(shell, "--noprofile", "--norc")

		cmdShell.Env = append(os.Environ(),
			"WINEPREFIX="+prefixPath,
			"PS1=[winebox:"+prefixName+"] \\w $ ",
			"HISTFILE="+historyFile,
		)

		cmdShell.Stdin = os.Stdin
		cmdShell.Stdout = os.Stdout
		cmdShell.Stderr = os.Stderr

		fmt.Println("Entering Wine shell:", prefixName)
		fmt.Printf("To exit Wine shell use 'exit'.\n\n")

		err := cmdShell.Run()

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == 127 {
					return
				}
			}

			fmt.Println("ERROR:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
