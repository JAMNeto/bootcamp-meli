package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrescentOrdering(t *testing.T) {
	nums := []int{4, 2, 3, 1}
	expectedResult := []int{1, 2, 3, 4}
	result := CrescentOrdering(nums)
	assert.Equal(t, expectedResult, result, "Devem ser iguais")
}
