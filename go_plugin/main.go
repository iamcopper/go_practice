package main

import "plugin"

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
}
