package string

func TestAllFilters( testElement interface{}, value string ) bool {

	switch converted := testElement.(type) {

	case nil:			  return true
	case Equal:           return converted.Test( value )
	case NotEqual:        return converted.Test( value )
	case HasPrefix:       return converted.Test( value )
	case NotHasPrefix:    return converted.Test( value )
	case HasSuffix:       return converted.Test( value )
	case NotHasSuffix:    return converted.Test( value )
	case Contains:        return converted.Test( value )
	case NotContains:     return converted.Test( value )
	case ContainsAny:     return converted.Test( value )
	case NotContainsAny:  return converted.Test( value )
	case InList:          return converted.Test( value )
	case NotInList:       return converted.Test( value )

	}

	return false
}
