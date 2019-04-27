package numeric

type GreaterThanOrEqual int64

func ( el GreaterThanOrEqual ) Test( value int64 ) bool {
	return value >= int64( el )
}
