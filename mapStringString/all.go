package mapStringString

func TestAllFilters( testElement interface{}, value interface{} ) bool {

	switch converted := testElement.(type) {

	case nil:				  return true
	case ContainsKey:         return converted.Test( value.(string) )
	case ContainsValue:       return converted.Test( value.(string) )
	case HasKeyInList:        return converted.Test( value.([]string) )
	case HasValueInList:      return converted.Test( value.([]string) )
	case KeyEqual:            return converted.Test( value.(string) )
	case KeyHasPrefix:        return converted.Test( value.(string) )
	case KeyHasSuffix:        return converted.Test( value.(string) )
	case KeyNotEqual:         return converted.Test( value.(string) )
	case KeyNotHasPrefix:     return converted.Test( value.(string) )
	case KeyNotHasSuffix:     return converted.Test( value.(string) )
	case NotContainsKey:      return converted.Test( value.(string) )
	case NotContainsValue:    return converted.Test( value.(string) )
	case NotHasKeyInList:     return converted.Test( value.([]string) )
	case NotHasValueInList:   return converted.Test( value.([]string) )
	case ValueEqual:          return converted.Test( value.(string) )
	case ValueHasPrefix:      return converted.Test( value.(string) )
	case ValueHasSuffix:      return converted.Test( value.(string) )

	}

	return false
}
