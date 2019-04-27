package numeric

type InClosedBetween    []int64

func ( el InClosedBetween ) Test( value int64 ) bool {
	converted := []int64( el )
	if value >= converted[0] && value <= converted[1] {
		return true
	}

	return false
}
