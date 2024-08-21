package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var empty bool = true

	carry := 0

	for {
		result := l1.Val + l2.Val + carry
		sum := result % 10
		var newNode ListNode = ListNode{Val: sum}
		carry = result / 10

		if empty {
			head = &newNode
			tail = &newNode
			empty = false
		} else {
			tail.Next = &newNode
			tail = &newNode
		}

		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = &ListNode{Val: 0}
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &ListNode{Val: 0}
		}
	}
	return head
}
