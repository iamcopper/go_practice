package main

import "fmt"

//定义一种函数类型
type Caculate func(a, b int) int

func add(a, b int) int {
	return a + b
}

func main() {
	//声明函数类型变量, 并将add函数赋值给函数类型变量
	//注意类型匹配
	var caculate Caculate = add
	//caculate := add

	fmt.Printf("add(1, 2) = %d\n", add(1, 2))
	fmt.Printf("caculate(1, 2) = %d\n", caculate(1, 2))

	fmt.Printf("\n\"add\" type: %T\n", add)
	fmt.Printf("\"caculate\" type: %T\n", caculate)
}
