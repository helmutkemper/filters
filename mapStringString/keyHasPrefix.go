package mapStringString

import "strings"

type KeyHasPrefix map[string]string

func ( el KeyHasPrefix ) Test( value string ) bool {
	for k := range el {
		if strings.HasPrefix(k, value) {
			return true
		}
	}

	return false
}