package arrayString

type NotHasAnyInList    []string

func ( el NotHasAnyInList ) Test( value []string ) bool {
	for _, elValue := range el {
		for _, valueInList := range value {
			if elValue == valueInList {
				return false
			}
		}
	}

	return true
}
