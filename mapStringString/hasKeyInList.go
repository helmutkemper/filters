package mapStringString

type HasKeyInList map[string]string

func ( el HasKeyInList ) Test( list []string ) bool {
	for k := range el {
		for _, valueInList := range list {
			if k == valueInList {
				return true
			}
		}
	}

	return false
}