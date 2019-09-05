package main

type Integer interface {
	Add(int)
	Value() int
	getInstant() *integer
}

type integer int

func NewInteger() Integer {
	i := new(integer)
	*i = 5
	return i
}

func (a *integer) Add(b int) {
	*a += integer(b)
}

func (a *integer) Value() int {
	return int(*a)
}

func (a *integer) getInstant() *integer {
	return a
}
