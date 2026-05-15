package prefix

import (
	"os"
	"path/filepath"
)

func Path(prefixName string) string {
	return filepath.Join(os.Getenv("HOME"), ".winebox", "prefixes", prefixName, "prefix")
}
