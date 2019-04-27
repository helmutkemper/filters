package numeric

type LessThanOrEqual int64

func ( el LessThanOrEqual ) Test( value int64 ) bool {
	return value <= int64( el )
}
