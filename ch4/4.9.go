package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		counts[in.Text()]++
	}

	for w, n := range counts {
		fmt.Printf("%q\t%d\n", w, n)
	}
}
