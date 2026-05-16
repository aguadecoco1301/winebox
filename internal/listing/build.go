/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package listing

import (
	"sort"

	"github.com/aguadecoco1301/winebox/internal/appstore"
	"github.com/aguadecoco1301/winebox/internal/prefix"
)

func Build() ([]Prefix, error) {
	prefixes, err := prefix.List()
	if err != nil {
		return nil, err
	}
	var result []Prefix
	for _, prefixName := range prefixes {
		store, err := appstore.Load(prefix.Path(prefixName))
		if err != nil {
			continue
		}

		listedPrefix := Prefix{
			Name:     prefixName,
			MainPath: store.Main,
		}

		var appNames []string

		for name := range store.Apps {
			appNames = append(appNames, name)
		}
		sort.Strings(appNames)

		for _, appName := range appNames {
			app := store.Apps[appName]
			listedPrefix.Apps = append(listedPrefix.Apps, App{
				Name: appName,
				Path: app.Path,
			})
		}
		result = append(result, listedPrefix)
	}
	return result, nil
}
