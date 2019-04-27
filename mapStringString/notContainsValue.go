package mapStringString

import "strings"

type NotContainsValue map[string]string

func ( el NotContainsValue ) Test( value string ) bool {
	for _, v := range el {
		if strings.Contains(v, value) {
			return false
		}
	}

	return true
}