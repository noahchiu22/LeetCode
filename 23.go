package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	curr := ans

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
		curr.Next = &ListNode{}
		curr = curr.Next

		fmt.Println("lists[mi]", lists[mi].Val, lists[mi].Next)
		lists[mi] = lists[mi].Next
	}

	return ans
}
