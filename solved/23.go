package main

import "fmt"

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	curr := ans

	if len(lists) == 0 || checkAllNil(lists) {
		return nil
	}

	for {
		var mi int = -1

		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if mi == -1 || lists[i].Val < lists[mi].Val {
				mi = i
			}
		}

		if mi == -1 {
			break
		}

		curr.Val = lists[mi].Val

		fmt.Println("lists[mi]", lists[mi].Val, lists[mi].Next)
		lists[mi] = lists[mi].Next

		if checkAllNil(lists) {
			break
		}

		curr.Next = &ListNode{}
		curr = curr.Next
	}

	return ans
}

func checkAllNil(lists []*ListNode) bool {
	for _, node := range lists {
		if node != nil {
			return false
		}
	}
	return true
}
