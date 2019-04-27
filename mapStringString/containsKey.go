package mapStringString

import "strings"

type ContainsKey map[string]string

func ( el ContainsKey ) Test( value string ) bool {
	for k := range el {
		if strings.Contains(k, value) {
			return true
		}
	}

	return false
}