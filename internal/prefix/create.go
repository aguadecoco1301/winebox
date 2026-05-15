/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package prefix

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

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
	prefixDir := filepath.Join(wineboxDir, "prefixes", name)

	// Verificar si existe
	_, err := os.Stat(filepath.Join(prefixDir, configFileName))
	if err == nil {
		return fmt.Errorf("prefix already exists")
	}

	err = os.MkdirAll(prefixDir, 0775)
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

	winePrefixDir := filepath.Join(prefixDir, "prefix")
	os.MkdirAll(winePrefixDir, 0775)
	wineCmd := exec.Command("wineboot")

	wineCmd.Env = append(
		os.Environ(),
		"WINEPREFIX="+winePrefixDir,
	)

	wineCmd.Stdout = os.Stdout //
	wineCmd.Stderr = os.Stderr // Para que wineboot no corra de forma silenciosa

	fmt.Println("Running:", wineCmd)
	err = wineCmd.Run()
	if err != nil {
		os.RemoveAll(prefixDir) // Rara vez funcionará. Solo cuando Wine cierre con un error, lo cual no suele ocurrir en un wineboot
		return err
	}

	fmt.Println("Created prefix:", name)
	return nil
}
