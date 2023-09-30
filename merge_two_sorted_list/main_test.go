package mergetwosortedlist

import "testing"

func TestMergeTwoSortedList(t *testing.T) {

	list1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	list2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	merged := mergeTwoLists(list1, list2)

	expected := []int{0, 2}
	for _, val := range expected {
		if merged.Val != val {
			t.Errorf("Expected %v receive %v", val, merged.Val)
		}
		merged = merged.Next
	}
}
