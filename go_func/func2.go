//可变参二次传递
package main

import (
	"errors"
	"fmt"
)

func main() {
	min, max, _ := minmax(7, 9, 3, 5, 1)
	fmt.Printf("min:%d, max:%d\n", min, max)

	arr := []int{7, 9, 3, 5, 1}
	min, max, _ = minmax(arr...)
	fmt.Printf("In the slice, min:%d, max:%d\n", min, max)
}

func minmax(a ...int) (min, max int, err error) {
	min, _ = minimum(a...)
	max, _ = maximum(a)
	return min, max, nil
}

func minimum(a ...int) (min int, err error) {
	if len(a) == 0 {
		return 0, errors.New("Parameter list is empty.")
	}

	min = a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min, nil
}

func maximum(a []int) (max int, err error) {
	if len(a) == 0 {
		return 0, errors.New("Parameter list is empty.")
	}

	max = a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	return max, nil
}
