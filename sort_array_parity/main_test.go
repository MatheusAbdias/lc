package sortarrayparity

import (
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	a := []int{3, 2, 5, 6}
	expected := []int{2, 6, 3, 5}
	output := sortArrayByParity(a)

	if !eq(output, expected) {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}

func eq(slice1 []int, slice2 []int) bool {
	for index, item := range slice1 {
		if slice2[index] != item {
			return false
		}
	}
	return true
}
