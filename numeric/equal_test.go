package numeric

import "fmt"

func ExampleEqual_Test() {
	var filter Equal = 5
	fmt.Printf( "%v\n", filter.Test( 5 ) )

	// output:
	// true
}
