//slice
package main

import (
	"fmt"
)

func main() {
	s := [10]int{1, 2, 3}
	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Printf("[out] s=%v\n", s)

	modifyArray1(s)
	fmt.Printf("[out] s=%v\n", s)

	modifyArray2(&s)
	fmt.Printf("[out] s=%v\n", s)
}

func modifyArray1(s [10]int) {
	s[0] = 11
	s[1] = 12
	s[2] = 13
	s[3] = 14
	fmt.Printf("[in]  s=%v\n", s)
}

func modifyArray2(s *[10]int) {
	(*s)[0] = 11
	(*s)[1] = 12
	(*s)[2] = 13
	(*s)[3] = 14
	fmt.Printf("[in]  *s=%v\n", *s)
}
