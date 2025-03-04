package main

func twoSum(nums []int, target int) (ans []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				ans = append(ans, i, j)
				return
			}
		}
	}
	return
}
