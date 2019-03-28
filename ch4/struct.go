package main

import "fmt"

func main() {
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes, X int
	}

	var w Wheel

	w.Circle.Point.X = 1
	w.Y = 2

	fmt.Printf("%#v\n", w)
}
