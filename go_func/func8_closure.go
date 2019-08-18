//闭包案例-函数计数器
package main

import "fmt"

func counter(f func()) func() int {
	n := 0
	return func() int {
		f()
		n += 1
		fmt.Printf("n = %d\n", n)
		return n
	}
}

//测试用的调用函数
func foo() {
	fmt.Println("foo is called")
}

func main() {
	cnt := counter(foo)

	cnt()
	cnt()
	cnt()
}
