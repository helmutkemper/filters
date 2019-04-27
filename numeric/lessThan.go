package numeric

type LessThan int64

func ( el LessThan ) Test( value int64 ) bool {
	return value < int64( el )
}
