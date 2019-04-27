package string

import "strings"

type NotHasPrefix    string

func ( el NotHasPrefix ) Test( value string ) bool {
	return !strings.HasPrefix( string( el ), value )
}
