package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	return dfs1448(root, root.Val)
}

func dfs1448(root *TreeNode, max int) (count int) {
	if root == nil {
		return 0
	}

	// if the current node is greater than or equal to the max, then it is a good node, and update max
	if root.Val >= max {
		max = root.Val
		count++
	}

	// recursive both leaves
	count += dfs1448(root.Left, max)
	count += dfs1448(root.Right, max)

	return count
}
