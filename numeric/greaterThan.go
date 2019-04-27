package numeric

type GreaterThan int64

func ( el GreaterThan ) Test( value int64 ) bool {
	return value > int64( el )
}
