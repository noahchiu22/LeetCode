package main

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		// Calculate area using min height * width
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}

		// Move pointer with smaller height inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
