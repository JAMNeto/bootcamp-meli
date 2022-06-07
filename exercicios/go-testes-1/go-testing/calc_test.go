package calc

import "testing"

func TestSum(t *testing.T) {
	num1 := 2
	num2 := 3
	expectedResult := 5
	result := Sum(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}
