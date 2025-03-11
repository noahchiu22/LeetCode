package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := &ListNode{}

	if list1 == nil && list2 == nil {
		return nil
	}

	temp := list3
	for {
		if list1 == nil {
			temp.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			temp.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			temp.Val = list1.Val
			list1 = list1.Next
		} else {
			temp.Val = list2.Val
			list2 = list2.Next
		}

		if list1 == nil && list2 == nil {
			break
		}

		temp.Next = &ListNode{}
		temp = temp.Next
	}

	return list3
}
