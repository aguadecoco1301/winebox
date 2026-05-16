/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package sanitize

import (
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Name(name string, allowUnsafe bool) (string, bool, error) {
	name = filepath.Clean(name)
	name = filepath.Base(name)

	if name == "." || name == ".." {
		return "", false, nil
	}

	// Remueve los acentos.
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	name, _, err := transform.String(t, name)
	if err != nil {
		return "", false, err
	}

	var suggestedName strings.Builder
	hasInvalidCharacters := false

	isAllowed := func(r rune) bool {
		return (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '-' || r == '.'
	}

	for _, char := range name {

		// espacios siempre se normalizan
		if char == ' ' {
			char = '-'
			hasInvalidCharacters = true
		}

		if isAllowed(char) {
			suggestedName.WriteRune(char)
		} else {
			hasInvalidCharacters = true

			// si NO es unsafe, se ignora silenciosamente en sanitización
			// si es unsafe, igual lo dejamos pasar pero avisamos
			if allowUnsafe {
				suggestedName.WriteRune(char)
			}
		}
	}

	return suggestedName.String(), hasInvalidCharacters, nil
}
