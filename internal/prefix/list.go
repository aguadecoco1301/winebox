/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package prefix

import (
	"os"
	"path/filepath"
)

func List() ([]string, error) {
	baseDir := filepath.Join(os.Getenv("HOME"), ".winebox", "prefixes")

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	var prefixes []string

	for _, entry := range entries {
		if entry.IsDir() {
			prefixes = append(prefixes, entry.Name())
		}
	}
	return prefixes, nil
}
