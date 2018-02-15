package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	result := Plus(10, 9)
	assert.Equal(t, result, int64(19))
}

func TestMinus(t *testing.T) {
	result := Minus(5, 10)
	assert.Equal(t, result, int64(-5))
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 9)
	assert.Equal(t, result, int64(45))
}

func TestDivide(t *testing.T) {
	result := Divide(10, 5)
	assert.Equal(t, result, int64(2))
}

func TestDivideByZero(t *testing.T) {
	result := Divide(10, 0)
	assert.Equal(t, result, int64(0))
}
