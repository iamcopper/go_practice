//传值&传引用: slice
package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3}

	fmt.Printf("Original: sl=%v\n", sl)

	appendSlice1(sl)
	fmt.Printf("After appendSlice1: sl=%v\n", sl)

	//appendSlice2(&sl)
	//fmt.Printf("After appendSlice2: sl=%v\n", sl)
}

//传递sclie
//append后，生成新的slice，但原slice(sl)并未改变，函数中也无法改变
func appendSlice1(s []int) {
	if s == nil {
		return
	}

	s = append(s, 11)
	fmt.Printf("In appendSlice1: s=%v\n", s)
}

//传递slice指针
//append后，生成新的sclice，在通过slice指针修改原slice(sl)
func appendSlice2(s *([]int)) {
	if s == nil {
		return
	}

	*s = append(*s, 12)
	fmt.Printf("In appendSlice1: *s=%v\n", *s)
}
