package main

import "fmt"

func main() {
	var sum int

	sum, _ = add(1, 2)
	fmt.Printf("sum = %d\n", sum)
}

func add(a, b int) (sum int, err error) {
	return a + b, nil
}
