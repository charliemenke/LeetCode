package main

import (
	mergetwosortedlists "github.com/charliemenke/LeetCode/merge-two-sorted-lists"
)

func main() {
	//fmt.Println(strconv.FormatBool(validparentheses.IsValid(("(])"))))

	list1 := mergetwosortedlists.ListNode{
		Val: 1,
		Next: &mergetwosortedlists.ListNode{
			Val: 2,
			Next: &mergetwosortedlists.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := mergetwosortedlists.ListNode{
		Val: 1,
		Next: &mergetwosortedlists.ListNode{
			Val: 3,
			Next: &mergetwosortedlists.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	_ = mergetwosortedlists.MergeTwoLists(&list1, &list2)
}
