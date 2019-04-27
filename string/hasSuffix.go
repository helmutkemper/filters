package string

import "strings"

type HasSuffix    string

func ( el HasSuffix ) Test( value string ) bool {
	return strings.HasSuffix( string( el ), value )
}
