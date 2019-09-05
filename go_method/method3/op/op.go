package op

type Integer interface {
	Add(int)
	sub(int)
	Value() int
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

func (a *integer) sub(b int) {
	*a -= integer(b)
}

func (a *integer) Value() int {
	return int(*a)
}
