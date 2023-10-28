package example

import "fmt"

type X struct {
	Field any
}

func Example() {
	var x *X
	fmt.Println(x.Field)
}
