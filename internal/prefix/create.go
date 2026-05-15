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
		Comment: "It is recommended to modify this file using Winebox instead of manual editing.",
	}

	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(prefixDir, configFileName), jsonData, 0664)
	if err != nil {
		return err
	}

	fmt.Println("Created prefix:", name)
	return nil
}
