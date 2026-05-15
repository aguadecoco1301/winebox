/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package appstore

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func filePath(prefixDir string) string {
	return filepath.Join(prefixDir, "apps.json")
}

func Load(prefixDir string) (Data, error) {
	var data Data

	path := filePath(prefixDir)

	file, err := os.Open(path)
	if err != nil {
		return Data{
			Apps: make(map[string]App),
		}, nil
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return data, err
	}

	if data.Apps == nil {
		data.Apps = make(map[string]App)
	}

	return data, nil
}

func Save(prefixDir string, data Data) error {
	path := filePath(prefixDir)
	tmp := path + ".tmp"

	file, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

// Core API

func AddApp(prefixDir, name, path string) error {
	data, err := Load(prefixDir)
	if err != nil {
		return err
	}

	data.Apps[name] = App{Path: path}

	return Save(prefixDir, data)
}

func SetMain(prefixDir, name string) error {
	data, err := Load(prefixDir)
	if err != nil {
		return err
	}

	data.Main = name

	return Save(prefixDir, data)
}

func DeleteApp(prefixDir, name string) error {
	data, err := Load(prefixDir)
	if err != nil {
		return err
	}

	delete(data.Apps, name)

	if data.Main == name {
		data.Main = ""
	}

	return Save(prefixDir, data)
}

func GetApp(prefixDir, name string) (App, bool, error) {
	data, err := Load(prefixDir)
	if err != nil {
		return App{}, false, err
	}

	app, ok := data.Apps[name]
	return app, ok, nil
}

func GetMainApp(prefixDir, name string) (App, bool, error) {
	data, err := Load(prefixDir)
	if err != nil {
		return App{}, false, err
	}

	if data.Main == "" {
		return App{}, false, nil
	}

	app, ok := data.Apps[data.Main]
	return app, ok, nil
}
