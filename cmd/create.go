/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aguadecoco1301/winebox/internal/prefix"
	"github.com/aguadecoco1301/winebox/internal/sanitize"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <prefix-name>",
	Short: "Creates a new WinePrefix",
	Long:  `Creates a new WinePrefix. The correct way to use it is to create a WinePrefix for each application you install.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]

		allowUnsafe, err := cmd.Flags().GetBool("allow-unsafe-name")
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		finalName, hasInvalidCharacters, err := sanitize.Name(prefixName, allowUnsafe)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if len(finalName) == 0 {
			fmt.Println("ERROR: empty or invalid prefix name")
			cmd.Help()
			return
		}

		// CASO 1: se usó allow-unsafe-name
		if allowUnsafe && hasInvalidCharacters {
			fmt.Println("WARNING: using unsafe prefix name. This may cause issues in Wine.")
		}

		// CASO 2: hubo caracteres invalidos (no se permite)
		if hasInvalidCharacters && !allowUnsafe {
			fmt.Println("ERROR: prefix name contains invalid characters")
			fmt.Println("Suggested safe name:", finalName)
			cmd.Help()
			return
		}

		fmt.Println("Creating new prefix:", finalName)

		err = prefix.CreatePrefix(finalName)
		if err != nil {
			fmt.Println("ERROR: failed to create prefix:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().BoolP("allow-unsafe-name", "a", false, "Allows any character in the prefix name. May cause errors with Wine.")
}
