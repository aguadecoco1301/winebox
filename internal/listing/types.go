/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package listing

type Prefix struct {
	Name     string
	MainPath string
	Apps     []App
	Size     uint64
}

type App struct {
	Name string
	Path string
}

type RenderOptions struct {
	ShowPaths bool
	ShowSize  bool
}

type TreeEntry struct {
	Label string
	Path  string
	Size  uint64
}
