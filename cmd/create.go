/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package cmd

import (
	"fmt"
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
	Use:   "create",
	Short: "Creates a new WinePrefix",
	Long:  `Creates a new WinePrefix. The correct way to use it is to create a WinePrefix for each application you install.`,
	Run: func(cmd *cobra.Command, args []string) {
		prefixName := strings.Join(args, " ")

		// Remueve los acentos.
		t := transform.Chain( //transform es que siga los pasos en ese orden. Función interesante
			norm.NFD,                           // Descompone los caracteres con sus diacríticos (ej. e + ´)
			runes.Remove(runes.In(unicode.Mn)), // Remueve los diacríticos
			norm.NFC,                           // Recompone los caracteres
		)
		prefixName, _, _ = transform.String(t, prefixName) // Aplica la transformación y descarta numero y error

		var suggestedName string
		var isUsedCorrectCharacters bool = true

		flags, err := cmd.Flags().GetBool("allow-unsafe-name")
		if err != nil {
			fmt.Println("Error: ", err)
		}

		if flags == false {
			isAllowedCharacters := func(r byte) bool {
				return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '.'
			}

			for i := 0; i < len(prefixName); i++ {
				char := prefixName[i]
				if char == ' ' {
					char = '-'
					isUsedCorrectCharacters = false
				}

				if isAllowedCharacters(char) {
					suggestedName += string(char)
				} else {
					isUsedCorrectCharacters = false
				}
			}
		} else {
			fmt.Println("WARNING: Using an unsafe name may cause Wine issues")
			suggestedName = prefixName
		}

		if isUsedCorrectCharacters {
			fmt.Println("New prefix: ", suggestedName) // Reemplazo de la creación del prefix
			err := prefix.CreatePrefix(suggestedName)
			if err != nil {
				fmt.Println("ERROR: Failed to create prefix: ", err)
			}
		} else {
			if suggestedName != "" && len(suggestedName) >= 2 {
				fmt.Println("ERROR: You entered an unsafe name. No action will be taken.\nIf you know what you are doing, check 'winebox create --help'.\nSuggested safe name:\n", suggestedName)
			} else {
				fmt.Println("ERROR: You entered an unsafe name. No action will be taken.\nIf you know what you are doing, check 'winebox create --help'.\nPlease use at least two of letters, numbers, '-' or '.'. ")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().BoolP("allow-unsafe-name", "a", false, "Allows any character in the prefix name. May cause errors with Wine.")
}
