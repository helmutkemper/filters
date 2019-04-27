package numeric

type NotEqual   int64

func ( el NotEqual ) Test( value int64 ) bool {
	return value != int64( el )
}
