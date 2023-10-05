package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	expected := []int{3}
	output := majorityElement(nums)

	if !eq(expected, output) {
		t.Errorf("Expected %v, receive %v", expected, output)
	}
}

func eq(x []int, y []int) bool {
	for index, v := range x {
		if v != y[index] {
			return false
		}
	}
	return true
}
