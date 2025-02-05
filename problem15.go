package main

import (
	"sort"
)

func threeSum(nums []int) (sets [][]int) {
	sort.Ints(nums)

	majorFindedInt := make(map[int]bool)

	for i := range nums {
		if majorFindedInt[nums[i]] {
			continue
		}

		minorFindedInt := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			if minorFindedInt[nums[j]] {
				continue
			}

			result := binarySearch(nums[j+1:], -(nums[i] + nums[j]))
			if result != -1 {
				sets = append(sets, []int{nums[i], nums[j], nums[j+1+result]})
			}

			minorFindedInt[nums[j]] = true
		}

		majorFindedInt[nums[i]] = true
	}

	return
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// func getThreeSum(nums []int, sets *[][]int, numMap map[string]bool) {
// 	if len(nums) < 3 || nums[0] > 0 {
// 		return
// 	}

// 	// pop the first element
// 	first := nums[0]
// 	nums = nums[1:]

// 	// find the two elements that add up to -first
// 	for i := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			x++
// 			key := strconv.Itoa(first) + strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
// 			if nums[i]+nums[j] == -first && !numMap[key] {
// 				*sets = append(*sets, []int{first, nums[i], nums[j]})
// 				numMap[key] = true
// 				if nums[i]+nums[j] == 0 {
// 					return
// 				}
// 			}
// 		}
// 	}

// 	getThreeSum(nums, sets, numMap)
// }
