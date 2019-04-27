package mapStringString

type KeyNotEqual map[string]string

func ( el KeyNotEqual ) Test( value string ) bool {
	for k := range el {
		if k == value {
			return false
		}
	}

	return true
}