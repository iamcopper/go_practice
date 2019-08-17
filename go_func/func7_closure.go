package main

import "fmt"

//type Fn func(int) int

func exfunc(base int) func(int) int {
	return func(x int) int {
		return (base + x)
	}
}

func main() {
	myfunc1 := exfunc(10)
	fmt.Printf("myfunc1(1)=%d, myfunc1(2)=%d\n", myfunc1(1), myfunc1(2))

	myfunc2 := exfunc(20)
	fmt.Printf("myfunc2(1)=%d, myfunc2(2)=%d\n", myfunc2(1), myfunc2(2))
}
