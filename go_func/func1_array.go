//array
package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println("[out]", arr)

	modifyArray(arr)
	fmt.Println("[out]", arr)
}

func modifyArray(a [10]int) {
	a[0] = 11
	a[1] = 12
	a[2] = 13
	a[3] = 14
	fmt.Println("[in] ", a)
}
