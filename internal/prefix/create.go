/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package prefix

import (
	"encoding/json"
	"fmt"
	"os"

	"path/filepath"

	"github.com/aguadecoco1301/winebox/internal/version"
)

var home, _ = os.UserHomeDir()
var wineboxDir = filepath.Join(home, ".winebox")

const configFileName = "config.json"

// https://chmod-calculator.com/
// Directorios:	0775 rwx | rwx | r-x
// Archivos: 	0664 rw- | rw- | r--

func CreatePrefix(name string) error {
	prefixDir := filepath.Join(wineboxDir, name)
	err := os.MkdirAll(prefixDir, 0775)
	if err != nil {
		return err
	}

	config := Config{
		Name:    name,
		Version: version.Version,
	}

	jsonData, err := json.MarshalIndent(config, "", "    ")

	os.WriteFile(filepath.Join(prefixDir, configFileName), jsonData, 0664)
	fmt.Println("Created prefix:", name)
	return nil
}
