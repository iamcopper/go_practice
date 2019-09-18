package main

import (
	"fmt"
)

func main() {
	sl2 := make([]int, 5)
	sl2[1] = 0
	fmt.Printf("sl2=%v, len(sl2)=%d, cap(sl2)=%d\n", sl2, len(sl2), cap(sl2))
	sl3 := sl2[2:4]
	fmt.Printf("sl3=%v, len(sl3)=%d, cap(sl3)=%d\n", sl3, len(sl3), cap(sl3))
}
