package searchinsertpos

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2

	expected := 1
	output := searchInsert(nums, target)

	if expected != output {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}
