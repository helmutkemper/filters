package string

type InList []string

func ( el InList ) Test( value string ) bool {
	for _, v := range el {
		if v == value {
			return true
		}
	}

	return false
}
