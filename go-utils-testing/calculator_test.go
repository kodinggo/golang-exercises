package goutilstesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := Tambah(1, 1)
		assert.Equal(t, 2, result)
	})

	t.Run("failed", func(t *testing.T) {
		result := Tambah(1, 1)
		assert.NotEqual(t, 3, result)
	})
}

func TestKurang(t *testing.T) {
	result := Kurang(3, 1)
	assert.Equal(t, 2, result)
}

func TestKali(t *testing.T) {
	result := Kali(2, 2)
	assert.Equal(t, 4, result)
}

func TestBagi(t *testing.T) {
	result := Bagi(8, 2)
	assert.Equal(t, 4, result)
}
