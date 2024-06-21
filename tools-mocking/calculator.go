package toolsmocking

import (
	"errors"
	"tools-mocking/utils"
)

type Calculator interface {
	Count(n1, n2 int, op rune) (int, error)
}

type calculator struct {
	math utils.Math
}

func (c calculator) Count(n1, n2 int, op rune) (result int, err error) {
	switch op {
	case '+':
		result = c.math.Add(n1, n2)
	case '-':
		// TODO
	case '*':
		// TODO
	case '/':
		// TODO
	default:
		err = errors.New("invalid operator")
	}

	return
}

func NewCalculator(math utils.Math) Calculator {
	return calculator{math: math}
}
