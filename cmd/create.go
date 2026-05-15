/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/aguadecoco1301/winebox/internal/prefix"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <prefix-name>",
	Short: "Creates a new WinePrefix",
	Long:  `Creates a new WinePrefix. The correct way to use it is to create a WinePrefix for each application you install.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := args[0]
		prefixName = filepath.Clean(prefixName)
		prefixName = filepath.Base(prefixName)

		if prefixName == "." || prefixName == ".." {
			fmt.Println("ERROR: invalid prefix name")
			cmd.Help()
			return
		}

		// Remueve los acentos.
		t := transform.Chain( //transform es que siga los pasos en ese orden. Función interesante
			norm.NFD,                           // Descompone los caracteres con sus diacríticos (ej. e + ´)
			runes.Remove(runes.In(unicode.Mn)), // Remueve los diacríticos
			norm.NFC,                           // Recompone los caracteres
		)
		prefixName, _, _ = transform.String(t, prefixName) // Aplica la transformación y descarta numero y error

		var suggestedName strings.Builder
		hasInvalidCharacters := false

		isAllowed := func(r rune) bool {
			return (r >= 'a' && r <= 'z') ||
				(r >= 'A' && r <= 'Z') ||
				(r >= '0' && r <= '9') ||
				r == '-' || r == '.'
		}

		allowUnsafe, err := cmd.Flags().GetBool("allow-unsafe-name")
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		for _, char := range prefixName {

			// espacios siempre se normalizan
			if char == ' ' {
				char = '-'
				hasInvalidCharacters = true
			}

			if isAllowed(char) {
				suggestedName.WriteRune(char)
			} else {
				hasInvalidCharacters = true

				// si NO es unsafe, se ignora silenciosamente en sanitización
				// si es unsafe, igual lo dejamos pasar pero avisamos
				if allowUnsafe {
					suggestedName.WriteRune(char)
				}
			}
		}

		finalName := suggestedName.String()

		if len(finalName) == 0 {
			fmt.Println("ERROR: empty or invalid prefix name")
			cmd.Help()
			return
		}

		// CASO 1: se usó allow-unsafe-name
		if allowUnsafe && hasInvalidCharacters {
			fmt.Println("WARNING: using unsafe prefix name. This may cause issues in Wine.")
		}

		// CASO 2: hubo sanitización (NO se permite)
		if hasInvalidCharacters && !allowUnsafe {
			fmt.Println("ERROR: prefix name contains invalid characters")
			fmt.Println("Suggested safe name:", finalName)
			cmd.Help()
			return
		}

		fmt.Println("Creating new prefix:", finalName)

		err = prefix.CreatePrefix(finalName)
		if err != nil {
			fmt.Println("ERROR: Failed to create prefix:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().BoolP("allow-unsafe-name", "a", false, "Allows any character in the prefix name. May cause errors with Wine.")
}
