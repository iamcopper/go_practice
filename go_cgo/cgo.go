// go call C with cgo
package main

import "fmt"

/*
#include <stdio.h>

void hi()
{
	printf("Hi, I'm in C\n");
}
*/
import "C"

func main() {
	C.hi()
	fmt.Println("Hi, I'm in go.");
}
