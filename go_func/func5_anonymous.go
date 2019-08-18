//匿名函数
package main

import "fmt"

func main() {
	func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Printf("sum = %d\n", sum)
	}()
}
