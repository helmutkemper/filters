package numeric

type NotInOpenedBetween []int64

func ( el NotInOpenedBetween ) Test( value int64 ) bool {
	converted := []int64( el )
	if value > converted[0] && value < converted[1] {
		return false
	}

	return true
}
