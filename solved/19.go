package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// count the length of the ListNode
	length := 1
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		length++
	}

	// target node
	targetNode := length - n + 1
	if targetNode == 1 {
		return head.Next
	}

	nthNode := 1
	temp = head
	for temp.Next != nil {
		if nthNode == targetNode-1 {
			// middle node
			temp.Next = temp.Next.Next
			fmt.Println("nthNode", nthNode)
			break
		}
		temp = temp.Next
		nthNode++
	}

	return head
}
