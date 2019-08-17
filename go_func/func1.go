package main

import "fmt"

type Caculate func(a, b int) int

func add(a, b int) int {
	return a + b
}

func main() {
	//caculate := add
	var caculate Caculate
	caculate = add

	fmt.Printf("add(1, 2) = %d\n", add(1, 2))
	fmt.Printf("caculate(1, 2) = %d\n", caculate(1, 2))

	//fmt.Printf("\n\"add\" type: %T\n", add)
	//fmt.Printf("\n\"caculate\" type: %T\n", caculate)
}
