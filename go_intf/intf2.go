package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	(*a) += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1

	fmt.Printf("a=%v\n", a)

	//值类型(a)只实现了值方法（Less），不满足LessAdder接口，
	//所以不能赋值给b
	//var b LessAdder = a   //会编译错误

	//指针类型(&a)既实现了值方法（Less）, 也实现了指针方法（Add），
	//满足LessAdder接口，所以可以赋值给b
	var b LessAdder = &a

	b.Add(2)
	fmt.Printf("After calling b.Add(2)=%v\n", a)
}
