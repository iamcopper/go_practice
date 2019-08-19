package main

import "fmt"

type Integer int

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) Sub(b Integer) {
	(*a) -= b
}

func main() {
	var a Integer = 5
	fmt.Printf("Original a = %v\n", a)

	a.Add(1)
	fmt.Printf("After a.Add(1), a = %v\n", a)
	(&a).Add(1)
	fmt.Printf("After (&a).Add(1), a = %v\n", a)

	a.Sub(1)
	fmt.Printf("After a.Sub(1), a = %v\n", a)
	(&a).Sub(1)
	fmt.Printf("After (&a).Sub(1), a = %v\n", a)
}
