package string

import "strings"

type ContainsAny    string

func ( el ContainsAny ) Test( value string ) bool {
	return strings.ContainsAny( string( el ), value )
}
