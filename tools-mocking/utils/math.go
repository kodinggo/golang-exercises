package utils

import "errors"

// Interface miliki object math
type Math interface {
	Add(n1, n2 int) int
	Subtract(n1, n2 int) (int, error)
}

// object math
type math struct {
}

func (m math) Add(n1, n2 int) int {
	return n1 + n2
}

func (m math) Subtract(n1, n2 int) (int, error) {
	if n1 < n2 {
		return 0, errors.New("n1 tidak boleh lebih kecil dari n2")
	}
	return n1 - n2, nil
}

func NewMath() Math {
	return math{}
}
