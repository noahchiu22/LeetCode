package main

func change(amount int, coins []int) int {
	// store num of methods of each amount
	dp := make([]int, amount+1)
	// 1 way to make amount 0
	dp[0] = 1

	// count each coin's num of methods of each amount
	for _, coin := range coins {
		// refresh num of methods for each amount
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
