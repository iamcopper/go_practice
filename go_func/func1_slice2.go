package main

import (
	"fmt"
)

type dataPackage struct {
	data    []uint8
	datalen uint32
}

func main() {
	var array [10]uint8
	sl := array[6:]

	pkg := new(dataPackage)
	sl = pkg.data

	fmt.Printf("out: array=%v, len=%d, cap=%d\n", array, len(array), cap(array))
	fmt.Printf("out: sl=%v, len=%d, cap=%d\n", sl, len(sl), cap(sl))

	plus(sl[3:])

	fmt.Printf("out: sl=%v, len=%d, cap=%d\n", sl, len(sl), cap(sl))
	fmt.Printf("out: array=%v, len=%d, cap=%d\n", array, len(array), cap(array))
}

func plus(s []uint8) {
	if s == nil {
		return
	}

	for i := 0; i < len(s); i++ {
		s[i] += 1
	}
	fmt.Printf("in: sl=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
