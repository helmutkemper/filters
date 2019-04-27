package string

import "strings"

type NotHasSuffix    string

func ( el NotHasSuffix ) Test( value string ) bool {
	return !strings.HasSuffix( string( el ), value )
}
