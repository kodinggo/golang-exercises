package toolsmocking

import (
	"testing"
	"tools-mocking/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	m := new(mocks.Math)
	c := NewCalculator(m)

	t.Run("success", func(t *testing.T) {
		m.On("Add", 2, 1).Return(2).Once()

		result, err := c.Count(1, 1, '+')
		assert.NoError(t, err)
		assert.Equal(t, 2, result)
	})
}
