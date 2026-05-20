/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package desktop

import "fmt"

func Build(entry Entry) string {
	return fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s
Type=Application
Terminal=false
StartupNotify=true
`, entry.Name, entry.Exec)
}

func FileName(prefixName string, appName string) string {
	internalName := prefixName
	if appName != "" {
		internalName += "-" + appName
	}

	return fmt.Sprintf("winebox-%s.desktop", internalName)
}
