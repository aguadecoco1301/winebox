/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package appstore

import (
	"encoding/json"
	"fmt"
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

func SetMain(prefixDir, path string) error {
	data, err := Load(prefixDir)
	if err != nil {
		return err
	}

	data.Main = path

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

func EditApp(prefixDir, oldName, newName, newPath string) error {
	data, err := Load(prefixDir)
	if err != nil {
		return err
	}

	app, ok := data.Apps[oldName]
	if !ok {
		return fmt.Errorf("app not found: %s", oldName)
	}

	delete(data.Apps, oldName)

	if newName == "" {
		newName = oldName
	}

	if newPath != "" {
		app.Path = newPath
	}

	data.Apps[newName] = app

	if data.Main == oldName {
		data.Main = newName
	}

	return Save(prefixDir, data)
}

func (d Data) GetMainApp() (App, error) {
	if d.Main == "" {
		return App{}, fmt.Errorf("no main app defined")
	}

	return App{Path: d.Main}, nil
}

func (d Data) Resolve(name string) (App, error) {
	if name == "" {
		return d.GetMainApp()
	}

	app, ok := d.Apps[name]
	if !ok {
		return App{}, fmt.Errorf("app not found: %s", name)
	}

	return app, nil
}
