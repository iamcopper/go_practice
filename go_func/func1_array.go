//slice
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

	modifyArray1(arr)
	fmt.Println("[out]", arr)

	modifyArray2(&arr)
	fmt.Println("[out]", arr)
}

func modifyArray1(a [10]int) {
	a[0] = 11
	a[1] = 12
	a[2] = 13
	a[3] = 14
	fmt.Println("[in] ", a)
}

func modifyArray2(aptr *[10]int) {
	(*aptr)[0] = 11
	(*aptr)[1] = 12
	(*aptr)[2] = 13
	(*aptr)[3] = 14
	fmt.Println("[in] ", *aptr)
}
