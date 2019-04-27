package string

type NotInList []string

func ( el NotInList ) Test( value string ) bool {
	for _, v := range el {
		if v == value {
			return false
		}
	}

	return true
}
