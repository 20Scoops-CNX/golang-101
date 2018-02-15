package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToIntSucccess(t *testing.T) {
	n1, n2, err := ConvertStringToInt("1", "2")
	assert.Equal(t, n1, int64(1))
	assert.Equal(t, n2, int64(2))
	assert.Nil(t, err)
}

func TestConvertStringToIntNumberOneError(t *testing.T) {
	n1, n2, err := ConvertStringToInt("d", "2")
	assert.Equal(t, n1, int64(0))
	assert.Equal(t, n2, int64(0))
	assert.NotNil(t, err)
}

func TestConvertStringToIntNumberTwoError(t *testing.T) {
	n1, n2, err := ConvertStringToInt("1", "aaaaa")
	assert.Equal(t, n1, int64(0))
	assert.Equal(t, n2, int64(0))
	assert.NotNil(t, err)
}

func TestGetNumberOperatorSuccess(t *testing.T) {
	result, err := GetNumberOperator("1")
	assert.Equal(t, result, int64(1))
	assert.Nil(t, err)
}

func TestGetNumberOperatorError(t *testing.T) {
	result, err := GetNumberOperator("a")
	assert.Equal(t, result, int64(0))
	assert.NotNil(t, err)
}

func TestGetResultPlus(t *testing.T) {
	result := GetResult(1, 2, 2)
	assert.Equal(t, result, int64(4))
}

func TestGetResultMinus(t *testing.T) {
	result := GetResult(2, 2, 2)
	assert.Equal(t, result, int64(0))
}
func TestGetResultMultiply(t *testing.T) {
	result := GetResult(3, 2, 2)
	assert.Equal(t, result, int64(4))
}

func TestGetResultDivide(t *testing.T) {
	result := GetResult(4, 2, 2)
	assert.Equal(t, result, int64(1))
}

func TestGetResultDivideByZero(t *testing.T) {
	result := GetResult(4, 2, 0)
	assert.Equal(t, result, int64(0))
}
