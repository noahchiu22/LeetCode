package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}

	var dp func(int, int, []int)
	dp = func(i, times int, existNum []int) {
		if i > len(candidates)-1 {
			return
		}

		// deep copy, because is same underlying array when length is same
		newExistNum := make([]int, len(existNum))
		copy(newExistNum, existNum)

		num := 0
		for j := 0; j < times; j++ {
			num += candidates[i]
		}
		for _, v := range newExistNum {
			num += v
		}
		fmt.Println("i", i, "num", num)

		if num > target {
			return
		}

		if num == target {
			for j := 0; j < times; j++ {
				newExistNum = append(newExistNum, candidates[i])
			}
			ans = append(ans, newExistNum)
			fmt.Println("ans", ans)
			return
		}

		dp(i, times+1, newExistNum)
		for j := 0; j < times; j++ {
			newExistNum = append(newExistNum, candidates[i])
		}
		dp(i+1, 0, newExistNum)
	}

	dp(0, 0, []int{})

	return ans
}
