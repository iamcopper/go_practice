package main

import (
	"fmt"
)

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 0}
	fmt.Printf("sl=%v, len(sl)=%d\n", sl, len(sl))

	plus(sl[3:6])
	fmt.Printf("sl=%v, len(sl)=%d\n", sl, len(sl))
}

func plus(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] += 5
	}
}
