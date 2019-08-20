package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3}
	index := 1

	fmt.Printf("sl=%v, index=%d\n", sl, index)

	plus(sl, index)
	fmt.Printf("sl=%v, index=%d\n", sl, index)
}

func plus(s []int, index int) {
	if s == nil {
		return
	}

	index += 1

	for i := 0; i < len(s); i++ {
		s[i] += index
	}
}
