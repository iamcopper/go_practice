package main

import (
	"fmt"
)

func main() {
	a := NewInteger()

	b := a.getInstant()

	a.Add(1)
	fmt.Printf("a.Value()=%v, *b=%v\n", a.Value(), *b)

	*b = 102
	fmt.Printf("a.Value()=%v, *b=%v\n", a.Value(), *b)
}
