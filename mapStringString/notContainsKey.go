package mapStringString

import "strings"

type NotContainsKey map[string]string

func ( el NotContainsKey ) Test( value string ) bool {
	for k := range el {
		if strings.Contains(k, value) {
			return false
		}
	}

	return true
}