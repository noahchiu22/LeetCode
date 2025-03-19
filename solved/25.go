package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head

	n := 1
	for curr != nil {
		if n%k != 0 && n%k < k/2+1 {
			curr = swap(curr, k-2*(n%k-1))
			if curr == nil {
				break
			}
		}
		curr = curr.Next
		n++
	}

	return head
}

func swap(node *ListNode, target int) *ListNode {
	curr := node
	s := 0
	for curr.Next != nil {
		s++
		if s == target-1 {
			node.Val, curr.Next.Val = curr.Next.Val, node.Val
			break
		}
		curr = curr.Next
	}
	if s < target-1 {
		return nil
	}
	return node
}
