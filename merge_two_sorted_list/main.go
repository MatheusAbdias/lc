package mergetwosortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Val = list1.Val
			result.Next = mergeTwoLists(list1.Next, list2)
		} else {
			result.Val = list2.Val
			result.Next = mergeTwoLists(list1, list2.Next)
		}
	} else {
		if list1 != nil {
			return list1
		}
		return list2
	}

	return result
}
