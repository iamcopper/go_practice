package main

import "fmt"

var a = 0

func init() {
	a += 1
}

func init() {
	a += 1
}

func main() {
	fmt.Println("a = ", a)
}
