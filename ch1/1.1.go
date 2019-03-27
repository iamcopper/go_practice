package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seq string
	for _, arg := range os.Args {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}
