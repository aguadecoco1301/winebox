/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>

*/
/*
winebox list
	gta-sa
	├─ gta-sa [main]
	├─ samp
	└─ map

	photoshop
	├─ photoshop [main]
	└─ updater

winebox list --size
	gta-sa (4.2G)
	├─ gta-sa [main]
	├─ samp
	└─ map

	photoshop (9.1G)
	├─ photoshop [main]
	└─ updater

winebox list --paths
	gta-sa
	├─ gta-sa [main]
	│  └─ /home/adriel/Games/gta-sa.exe
	├─ samp
	│  └─ /home/adriel/Games/samp.exe
	└─ map
	   └─ /home/adriel/Games/map.exe

	osu
	└─ osu [main]
	   └─ /home/adriel/Games/osu.exe

winebox list --paths --size
	gta-sa (4.2G)
	├─ gta-sa [main]
	│  └─ /home/adriel/Games/gta-sa.exe
	├─ samp
	│  └─ /home/adriel/Games/samp.exe
	└─ map
	   └─ /home/adriel/Games/map.exe

	osu (1.1G)
		└─ osu [main]
			└─ /home/adriel/Games/osu.exe

winebox list --prefixes
	gta-sa
	photoshop
	osu

winebox list gta-sa
	NAME       MAIN   PATH
	gta-sa     yes    /home/adriel/Games/gta-sa.exe
	samp       no     /home/adriel/Games/samp.exe
	map        no     /home/adriel/Games/map.exe

winebox list gta-sa --size
	gta-sa      4.2G
	osu         1.1G
	photoshop   9.1G
*/
package cmd

import (
	"fmt"

	"github.com/aguadecoco1301/winebox/internal/listing"
	"github.com/aguadecoco1301/winebox/internal/prefix"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [prefix]",
	Short: "List Wine prefixes and registered applications",
	Long: `Displays Wine prefixes and their registered applications.

By default, Winebox shows all prefixes in a tree-style view similar to
tools like lsblk or tree.

You can enable additional information using flags such as:
  --size      Show prefix size on disk
  --paths     Show executable paths
  --prefixes  Show only prefix names

When a prefix name is provided, Winebox displays a detailed table view
of the applications registered inside that prefix.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := listing.Build()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		showPaths, err := cmd.Flags().GetBool("paths")
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		showSize, err := cmd.Flags().GetBool("size")
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		showPrefixes, err := cmd.Flags().GetBool("prefixes")
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if showPrefixes {
			listing.RenderPrefixes(data)
			return
		}

		if showSize {
			for i := range data {
				p, err := prefix.Size(prefix.Path(data[i].Name))
				if err != nil {
					continue
				}
				data[i].Size = p
			}
		}
		listing.RenderTree(data, listing.RenderOptions{
			ShowPaths: showPaths,
			ShowSize:  showSize,
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("paths", "p", false, "Show executable paths")
	listCmd.Flags().Bool("prefixes", false, "Show only prefix names")
	listCmd.Flags().BoolP("size", "s", false, "Show prefix size on disk")
}
