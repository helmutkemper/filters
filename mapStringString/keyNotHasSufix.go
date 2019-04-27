package mapStringString

import "strings"

type KeyNotHasSuffix map[string]string

func ( el KeyNotHasSuffix ) Test( value string ) bool {
	for k := range el {
		if strings.HasSuffix(k, value) {
			return false
		}
	}

	return true
}