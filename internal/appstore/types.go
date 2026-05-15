/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package appstore

type App struct {
	Path string `json:"path"`
}

type Data struct {
	Main string         `json:"main"`
	Apps map[string]App `json:"apps"`
}
