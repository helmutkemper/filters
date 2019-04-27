package mapStringString

import "strings"

type ContainsValue map[string]string

func ( el ContainsValue ) Test( value string ) bool {
	for _, v := range el {
		if strings.Contains(v, value) {
			return true
		}
	}

	return false
}