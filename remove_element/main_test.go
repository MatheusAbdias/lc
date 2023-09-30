package removeelement

import "testing"

func TestRemoveElemet(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expected := 5
	output := removeElement(nums, val)

	if output != expected {
		t.Errorf("Expected %v but receive %v", expected, output)
	}

}
