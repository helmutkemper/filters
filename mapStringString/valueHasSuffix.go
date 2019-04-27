package mapStringString

import "strings"

type ValueHasSuffix map[string]string

func ( el ValueHasSuffix ) Test( value string ) bool {
	for _, v := range el {
		if strings.HasSuffix(v, value) {
			return true
		}
	}

	return false
}