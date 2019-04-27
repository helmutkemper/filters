package numeric

// English: int64 filter equality
type Equal       int64

// English: int64 filter equality test
func ( el Equal ) Test( value int64 ) bool {
	return value == int64( el )
}