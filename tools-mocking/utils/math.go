package utils

type Math interface {
	Add(n1, n2 int) int
}

type math struct {
}

func (m math) Add(n1, n2 int) int {
	return n1 + n2
}

func NewMath() Math {
	return math{}
}
