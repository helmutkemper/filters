package numeric

func TestAllFilters( testElement interface{}, value int64 ) bool {

	switch converted := testElement.(type) {

	case nil:				  return true
	case Equal:               return converted.Test( value )
	case NotEqual:            return converted.Test( value )
	case InList:              return converted.Test( value )
	case NotInList:           return converted.Test( value )
	case InClosedBetween:     return converted.Test( value )
	case NotInClosedBetween:  return converted.Test( value )
	case InOpenedBetween:     return converted.Test( value )
	case NotInOpenedBetween:  return converted.Test( value )
	case LessThan:            return converted.Test( value )
	case LessThanOrEqual:     return converted.Test( value )
	case GreaterThan:         return converted.Test( value )
	case GreaterThanOrEqual:  return converted.Test( value )

	}

	return false
}
