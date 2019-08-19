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

	a := s[:3]
	fmt.Printf("[main]\ns=%v len(s)=%d, cap(s)=%d\na=%v, len(a)=%d, cap(a)=%d\n\n", s, len(s), cap(s), a, len(a), cap(a))

	appendSlice1(a)
	fmt.Printf("[main]after appendSlice1:\ns=%v, a=%v\n\n", s, a)

	appendSlice2(&a)
	fmt.Printf("[main]after appendSlice2:\ns=%v, a=%v\n\n", s, a)

	/*
		b := s[]int{1, 2, 3}
		b = appendSlice3(b)
		fmt.Printf("After appendSlice3: b=%v\n\n", b)
	*/
}

//传递sclie
//append后，生成新的slice，原slice(实参a)并未改变，函数中也无法改变实参a
func appendSlice1(sl []int) {
	if sl == nil {
		return
	}

	sl = append(sl, 11)
	fmt.Printf("[appendSlice1]\nsl=%v\n\n", sl)
}

//传递slice指针
//append后，生成新的sclice，在通过slice指针修改原slice(实参a)
func appendSlice2(slptr *[]int) {
	if slptr == nil {
		return
	}

	*slptr = append(*slptr, 12)
	fmt.Printf("[appendSlice2]\n*slptr=%v\n\n", *slptr)
}

func appendSlice3(sl []int) []int {
	if sl == nil {
		return sl
	}

	sl = append(sl, 13)
	fmt.Printf("[appendSlice3]\nsl=%v\n\n", sl)
	return sl
}
