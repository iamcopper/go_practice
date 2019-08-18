package main

import "fmt"

func main() {
	fmt.Println("a open")
	defer fmt.Println("a close")

	fmt.Println("b open")
	defer fmt.Println("b close")

	fmt.Println("c open")
	defer fmt.Println("c close")
}
