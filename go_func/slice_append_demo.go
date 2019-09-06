package main

import "fmt"

func main() {
	//s := make([]int, 8)
	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	a := s[:3]
	fmt.Printf("[out] s=%v, a=%v\n", s, a)
	appendSlice(a)
	fmt.Printf("[out] s=%v, a=%v\n", s, a)
}

func appendSlice(s []int) {
	s = append(s, 11)
	fmt.Printf("[in]  s=%v\n", s)
}
