package string

import "strings"

type NotContains    string

func ( el NotContains ) Test( value string ) bool {
	return !strings.Contains( string( el ), value )
}
