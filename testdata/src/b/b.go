package b

type X struct {
	Field any
}

func f() {
	var x *X
	print(x.Field) // want "(nilness)"
}
