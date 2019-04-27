package mapStringString

import "strings"

type KeyNotHasPrefix map[string]string

func ( el KeyNotHasPrefix ) Test( value string ) bool {
	for k := range el {
		if strings.HasPrefix(k, value) {
			return false
		}
	}

	return true
}