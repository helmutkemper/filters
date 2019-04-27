package numeric

type InOpenedBetween    []int64

func ( el InOpenedBetween ) Test( value int64 ) bool {
	converted := []int64( el )
	if value > converted[0] && value < converted[1] {
		return true
	}

	return false
}
