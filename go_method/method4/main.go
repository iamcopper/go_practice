package main

import (
	"fmt"
)

func main() {
	a := NewInteger()

	// Define a method to get the interface instant.
	// To get the interface instant, so that we can use this instant do more things.
	b := a.getInstant()

	a.Add(1)
	fmt.Printf("a.Value()=%v, *b=%v\n", a.Value(), *b)

	*b = 102
	fmt.Printf("a.Value()=%v, *b=%v\n", a.Value(), *b)
}
