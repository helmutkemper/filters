package numeric

type NotInList  []int64

func ( el NotInList ) Test( value int64 ) bool {
	for _, v := range el {
		if v == value {
			return false
		}
	}
	return true
}
