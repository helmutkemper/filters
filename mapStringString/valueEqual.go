package mapStringString

type ValueEqual map[string]string

func ( el ValueEqual ) Test( value string ) bool {
	for _, v := range el {
		if v == value {
			return true
		}
	}

	return false
}