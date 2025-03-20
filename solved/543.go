package main

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := calcDiameter(root)
	return diameter
}

func calcDiameter(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	left, leftMax := calcDiameter(node.Left)
	right, rightMax := calcDiameter(node.Right)

	return 1 + int(math.Max(float64(left), float64(right))), int(math.Max(float64(left+right), math.Max(float64(leftMax), float64(rightMax))))
}
