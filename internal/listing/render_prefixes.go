/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package listing

import "fmt"

func RenderPrefixes(prefixes []Prefix) {
	for _, prefix := range prefixes {
		fmt.Println(prefix.Name)
	}
}
