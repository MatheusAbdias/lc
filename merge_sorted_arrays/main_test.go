package mergesortedarrays

import (
	"testing"
)

func TestMegerSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	exptected := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)

	if len(exptected) != len(nums1) {
		t.Errorf("Expected %v,receive %v", exptected, nums1)
	}
}
