/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package prefix

import (
	"os"
	"path/filepath"
)

func Size(path string) (uint64, error) {
	var total uint64

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += uint64(info.Size())
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return total, nil
}
