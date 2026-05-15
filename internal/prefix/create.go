/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package prefix

import (
	"fmt"
	"os"

	"path/filepath"
)

var home, _ = os.UserHomeDir()
var wineboxDir = filepath.Join(home, ".winebox")

// https://chmod-calculator.com/
// Directorios:	0775 rwx | rwx | r-x
// Archivos: 	0664 rw- | rw- | r--

func CreatePrefix(name string) error {
	prefixDir := filepath.Join(wineboxDir, name)
	err := os.MkdirAll(prefixDir, 0775)
	if err != nil {
		return err
	}
	os.WriteFile(filepath.Join(prefixDir, "config.json"), []byte(name), 0664)
	fmt.Println("Created prefix:", name)
	return nil
}
