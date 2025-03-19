package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head

	// node number
	n := 1
	for curr != nil {
		// only swap the nodes are less than half of the group
		if n%k != 0 && n%k < k/2+1 {
			// swap the nodes
			// k-2*(n%k-1) is the target node of the group
			// if n%k = 1, means n the first node of the group
			// so the target node of the group is k-2*(1-1) = k
			// means curr to target node takes k steps
			// if n%k = 2, means n the second node of the group
			// so the target node of the group is k-2*(2-1) = k-2
			// means curr to target node takes k-2 steps
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
	nextNode := 0
	for curr.Next != nil {
		nextNode++
		// if the next node is the target node, swap the nodes
		if nextNode == target-1 {
			node.Val, curr.Next.Val = curr.Next.Val, node.Val
			break
		}
		curr = curr.Next
	}
	// if the next node is not the target node, return nil
	if nextNode < target-1 {
		return nil
	}
	return node
}
