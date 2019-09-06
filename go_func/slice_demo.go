package main

import (
	"fmt"
)

func main() {
	array := [3]int{1, 2, 3}
	fmt.Printf("array=%v, len(array)=%d, cap(array)=%d\n", array, len(array), cap(array))
	sl1 := array[1:1]
	fmt.Printf("sl1=%v, len(sl1)=%d, cap(sl1)=%d\n", sl1, len(sl1), cap(sl1))
	sl1 = array[2:3]
	fmt.Printf("sl1=%v, len(sl1)=%d, cap(sl1)=%d\n", sl1, len(sl1), cap(sl1))

	sl2 := make([]int, 5)
	fmt.Printf("sl2=%v, len(sl2)=%d, cap(sl2)=%d\n", sl2, len(sl2), cap(sl2))
	sl2[0] = 0
	sl2[1] = 0
	fmt.Printf("sl2=%v, len(sl2)=%d, cap(sl2)=%d\n", sl2, len(sl2), cap(sl2))
}
