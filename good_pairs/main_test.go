package goodpairs

import "testing"

func TestGoodPairs(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	expected := 4
	output := numIdenticalPairs(nums)

	if output != expected {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}
