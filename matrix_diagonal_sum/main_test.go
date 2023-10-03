package matrixdiagonalsum

import "testing"

func TestDiagonalSum(t *testing.T) {
	mat := [][]int{
		{7, 3, 1, 9},
		{3, 4, 6, 9},
		{6, 9, 6, 6},
		{9, 5, 8, 5},
	}

	expected := 55
	output := diagonalSum(mat)

	if output != expected {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}
