package numeric

type InList     []int64

func ( el InList ) Test( value int64 ) bool {
	for _, v := range el {
		if v == value {
			return true
		}
	}
	return false
}
