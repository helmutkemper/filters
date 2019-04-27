package mapStringString

type KeyEqual map[string]string

func ( el KeyEqual ) Test( value string ) bool {
	for k := range el {
		if k == value {
			return true
		}
	}

	return false
}