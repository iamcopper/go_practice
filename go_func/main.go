package main

import "fmt"

func main() {
	var sum int

	sum, _ = add(1, 2)
	fmt.Printf("sum = %d\n", sum)
}

//匿名返回值
func add(a, b int) (int, error) {
	return a + b, nil
}

/*
//命名返回值
func add(a, b int) (sum int, err error) {
	sum = a + b
	err = nil
	return
}
*/
