package string

import "strings"

type NotEqual        string

func ( el NotEqual ) Test( value string ) bool {
	return strings.Compare( string( el ), value ) != 0
}
