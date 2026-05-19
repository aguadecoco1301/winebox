/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package desktop

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(fileName string, content string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	desktopPath := filepath.Join(home, ".local", "share", "applications", fileName)

	err = os.MkdirAll(filepath.Dir(desktopPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(desktopPath, []byte(content), 0644)
	if err != nil {
		return err
	}

	os.Chmod(desktopPath, 0755)

	fmt.Println("Desktop entry generated:", desktopPath)
	return nil
}
