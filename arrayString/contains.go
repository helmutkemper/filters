package arrayString

import "strings"

type Contains    []string

func ( el Contains ) Test( value string ) bool {
	for _, v := range el {
		if strings.Contains( v, value ) {
			return true
		}
	}

	return false
}
