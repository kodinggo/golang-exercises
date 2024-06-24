package toolsmocking

import (
	"errors"
	"testing"
	"tools-mocking/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	mathMock := new(mocks.Math)
	c := NewCalculator(mathMock)

	t.Run("success when add process", func(t *testing.T) {
		// Memastikan method math.Add dipanggil dengan argument yang benar
		mathMock.On("Add", 1, 1).Return(2).Once()

		result, err := c.Count(1, 1, '+')

		assert.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("success when subtract process", func(t *testing.T) {
		// Memastikan method math.Subtract dipanggil dengan argument yang benar
		mathMock.On("Subtract", 10, 5).Return(5, nil).Once()

		result, err := c.Count(10, 5, '-')

		assert.NoError(t, err)
		assert.Equal(t, 5, result)
	})

	t.Run("error when subtract process", func(t *testing.T) {
		// Memastikan method math.Subtract dipanggil dengan argument yang benar
		mathMock.On("Subtract", 10, 5).Return(5, errors.New("terjadi error")).Once()

		result, err := c.Count(10, 5, '-')

		assert.Error(t, err)
		assert.Equal(t, 5, result)
	})

	t.Run("error invalid operator", func(t *testing.T) {
		result, err := c.Count(10, 5, '^')

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}
