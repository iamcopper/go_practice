package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("testplugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	/* Interface type assert */
	hello, ok := f.(func())
	if !ok {
		panic("not found")
	}

	hello()

	/* Using type switch to check the interface type */
	switch t := f.(type) {
	case func():
		fmt.Printf("It's a func() type, type=%T value=%v\n", t, t)
	case int:
		fmt.Printf("It's a int type, type=%T value=%v\n", t, t)
	default:
		fmt.Printf("It's a unknown type, type=%T, value=%v\n", t, t)
	}

}
