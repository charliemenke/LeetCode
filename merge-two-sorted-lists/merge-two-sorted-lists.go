package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	toReturn := ListNode{Val: 0, Next: nil}
	toReturnPointer := &toReturn
	list1Pointer := list1
	list2Pointer := list2
	if list1 != nil {
		for {
			if (list1Pointer != nil && list2Pointer != nil) && list1Pointer.Val <= list2Pointer.Val {
				toReturnPointer.Val = list1Pointer.Val
				list1Pointer = list1Pointer.Next
			} else if (list1Pointer != nil && list2Pointer != nil) && list1Pointer.Val > list2Pointer.Val {
				toReturnPointer.Val = list2Pointer.Val
				list2Pointer = list2Pointer.Next
			} else {
				if list1Pointer == nil {
					toReturnPointer.Val = list2Pointer.Val
					list2Pointer = list2Pointer.Next
				} else {
					toReturnPointer.Val = list1Pointer.Val
					list1Pointer = list1Pointer.Next
				}
			}
			if list1Pointer == nil && list2Pointer == nil {
				break
			} else {
				toReturnPointer.Next = new(ListNode)
				toReturnPointer = toReturnPointer.Next
			}

		}
	} else {
		return list2
	}
	return &toReturn
}
