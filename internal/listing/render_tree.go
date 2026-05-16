/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package listing

import "fmt"

func RenderTree(prefixes []Prefix, opts RenderOptions) {
	for _, prefix := range prefixes {
		fmt.Print(prefix.Name)
		if opts.ShowSize {
			fmt.Printf(" (%s)", HumanSize(prefix.Size))
		}
		fmt.Println()
		// Build entries
		var entries []TreeEntry
		if prefix.MainPath != "" {
			entries = append(entries, TreeEntry{
				Label: prefix.Name + " [main]",
				Path:  prefix.MainPath,
			})
		}

		for _, app := range prefix.Apps {
			entries = append(entries, TreeEntry{
				Label: app.Name,
				Path:  app.Path,
			})
		}

		// Render entries
		for i, entry := range entries {
			treePrefix := "├─"
			last := i == len(entries)-1
			if last {
				treePrefix = "└─"
			}
			fmt.Printf("%s %s\n", treePrefix, entry.Label)

			if opts.ShowPaths {
				pathPrefix := "│"

				if last {
					pathPrefix = " "
				}
				fmt.Printf("%s  └─ %s\n", pathPrefix, entry.Path)
			}
		}
		fmt.Println()
	}
}
