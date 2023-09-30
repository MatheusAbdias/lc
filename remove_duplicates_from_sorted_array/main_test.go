package removeduplicatesfromsortedarray

import (
	"testing"
)

func TestRemoveDubplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// [0,1,2,3,4]
	expected := 5

	value := removeDuplicates(nums)

	if value != expected {
		t.Errorf("Expected %v receive %v", expected, value)
	}
}
