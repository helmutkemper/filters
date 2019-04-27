package string

import "strings"

type Contains    string

func ( el Contains ) Test( value string ) bool {
	return strings.Contains( string( el ), value )
}
