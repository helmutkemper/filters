package arrayString

type NotHasAllInList    []string

func ( el NotHasAllInList ) Test( value []string ) bool {
	count := 0
	for _, elValue := range el {
		for _, valueInList := range value {
			if elValue == valueInList {
				count += 1
			}
		}
	}

	return count != len( value )
}
