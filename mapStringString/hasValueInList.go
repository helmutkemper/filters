package mapStringString

type HasValueInList map[string]string

func ( el HasValueInList ) Test( value []string ) bool {
	for _, elValue := range el {
		for _, valueInList := range value {
			if elValue == valueInList {
				return true
			}
		}
	}

	return false
}