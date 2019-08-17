//通过函数值实现将一种行为参数化函数
package main

import "fmt"

type Op func(int) int

func main() {
	a := 1
	b := callop(add1, a)
	c := callop(add2, a)
	d := callop(add10, a)

	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)
}

func callop(op Op, x int) int {
	return op(x)
}

func add1(x int) int {
	return x + 1
}

func add2(x int) int {
	return x + 2
}

func add10(x int) int {
	return x + 10
}
