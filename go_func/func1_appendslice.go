//slice
package main

import (
	"fmt"
)

func main() {
	//s := [3]int{1, 2, 3}
	s := [10]int{1, 2, 3}
	s[0] = 1
	s[1] = 2
	s[2] = 3

	a := s[:3]

	appendSlice1(a)
	fmt.Printf("s=%v, a=%v, a=%p\n", s, a, a)

	appendSlice2(&a)
	fmt.Printf("s=%v, a=%v, a=%p\n", s, a, a)
}

func appendSlice1(sl []int) {
	if sl == nil {
		return
	}

	sl = append(sl, 11)
}

func appendSlice2(slptr *[]int) {
	if slptr == nil {
		return
	}

	*slptr = append(*slptr, 12)
}
