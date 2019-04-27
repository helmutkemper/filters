package string

import "strings"

type NotContainsAny    string

func ( el NotContainsAny ) Test( value string ) bool {
	return !strings.ContainsAny( string( el ), value )
}
