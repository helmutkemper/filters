package mapStringString

type NotHasKeyInList map[string]string

func ( el NotHasKeyInList ) Test( value []string ) bool {
	for k := range el {
		for _, valueInList := range value {
			if k == valueInList {
				return false
			}
		}
	}

	return true
}