package string

import "strings"

type Equal        string

func ( el Equal ) Test( value string ) bool {
	return strings.Compare( string( el ), value ) == 0
}
