package arrayString

import "strings"

type NotContains    []string

func ( el NotContains ) Test( value string ) bool {
	for _, v := range el {
		if strings.Contains( v, value ) {
			return false
		}
	}

	return true
}
