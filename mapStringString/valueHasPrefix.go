package mapStringString

import "strings"

type ValueHasPrefix map[string]string

func ( el ValueHasPrefix ) Test( value string ) bool {
	for _, v := range el {
		if strings.HasPrefix(v, value) {
			return true
		}
	}

	return false
}