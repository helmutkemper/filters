package mapStringString

import "strings"

type KeyHasSuffix map[string]string

func ( el KeyHasSuffix ) Test( value string ) bool {
	for k := range el {
		if strings.HasSuffix(k, value) {
			return true
		}
	}

	return false
}