package mapStringString

type NotHasValueInList map[string]string

func ( el NotHasValueInList ) Test( value []string ) bool {
	for _, elValue := range el {
		for _, valueInList := range value {
			if elValue == valueInList {
				return false
			}
		}
	}

	return true
}