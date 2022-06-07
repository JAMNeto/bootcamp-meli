package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	num1 := 10
	num2 := 2
	expectedResult := 5
	result, err := Divide(num1, num2)
	assert.Equal(t, expectedResult, result, err)
	num3 := 0
	result2, err2 := Divide(num1, num3)
	assert.EqualError(t, err2, "O denominador não pode ser menor ou igual a zero", result2)
}
