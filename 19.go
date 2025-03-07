package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// count the length of the ListNode
	length := 0
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		length++
	}

	// find the nth node from the end
	targetNode := length - n + 1
	nthNode := 1
	temp = head
	for temp != nil {
		if nthNode == targetNode {
			if nthNode == 1 {
				// first node
				temp = temp.Next
			} else if temp.Next != nil {
				// middle node
				temp.Next = temp.Next.Next
			} else {
				// last node
				temp = nil
				break
			}
			fmt.Println("nthNode", nthNode, temp.Val)
		}
		temp = temp.Next
		nthNode++
	}

	return head
}
