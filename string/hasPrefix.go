package string

import "strings"

type HasPrefix    string

func ( el HasPrefix ) Test( value string ) bool {
	return strings.HasPrefix( string( el ), value )
}
