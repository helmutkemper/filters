package arrayString

type HasAnyInList    []string

func ( el HasAnyInList ) Test( value []string ) bool {
	for _, elValue := range el {
		for _, valueInList := range value {
			if elValue == valueInList {
				return true
			}
		}
	}

	return false
}
