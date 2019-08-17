//可变参
package main

import (
	"errors"
	"fmt"
)

func main() {
	m, _ := minimum(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", m)

	arr := []int{7, 9, 3, 5, 1}
	m, _ = minimum(arr...)
	fmt.Printf("The minimum in the array is: %d\n", m)
}

func minimum(a ...int) (m int, err error) {
	if len(a) == 0 {
		return 0, errors.New("Parameter list is empty.")
	}

	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min, nil
}
