package main

import (
	"fmt"
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	fmt.Println("nums", nums)
	fmt.Println("target", target)

	closest := nums[0] + nums[1] + nums[2]

	for i := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			temp := closestBinarySearch(nums[j+1:], target-(nums[i]+nums[j]))
			sum := nums[i] + nums[j] + temp
			if math.Abs(float64(sum-target)) <= math.Abs(float64(closest-target)) {
				closest = sum
			}
		}
	}

	return closest
}

func closestBinarySearch(nums []int, target int) (closest int) {
	left := 0
	right := len(nums) - 1
	mid := 0

	for {
		mid = (left + right) / 2
		closest = nums[mid]
		nextEdge := 0

		if right-left <= 1 {
			if nums[right] < target || nums[right]-target < target-nums[left] {
				closest = nums[right]
			}
			break
		}

		if nums[mid] < target {
			left = mid + 1
			nextEdge = nums[left]
		} else {
			right = mid - 1
			nextEdge = nums[right]
		}

		// mid hasn't changed, so we need to check if the nextEdge is closer to target
		leftEdge, rightEdge := nextEdge, nums[mid]

		if nextEdge > nums[mid] {
			leftEdge = nums[mid]
			rightEdge = nextEdge
		}

		// if nums[mid] is between leftEdge and rightEdge, then closest is one of them
		if leftEdge <= target && rightEdge >= target {
			if rightEdge-target > target-leftEdge {
				closest = leftEdge
			} else {
				closest = rightEdge
			}
			break
		}
	}

	return closest
}
