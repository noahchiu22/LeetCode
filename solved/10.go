package main

func isMatch(s string, p string) bool {
	// memoization for indices
	memo := make(map[[2]int]bool)
	// declare dp function first for nested function
	var dp func(i, j int) bool

	dp = func(i, j int) bool {
		indicesKey := [2]int{i, j}
		ans := false
		// only execute if not memoized yet
		if _, exists := memo[indicesKey]; !exists {
			// if only j at the end, ans is false, else if i is also at the end, ans is true
			if j >= len(p) {
				ans = i >= len(s)
			} else {
				// define if this rune is match
				isRuneMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')

				// if p[j+1] is *, there are two cases:
				// 1. ignore this p[j] (i,j+2)
				// 2. if this rune is match, use p[j] to check next rune (i+1,j)
				if j < len(p)-1 && p[j+1] == '*' {
					// don't need to parenthesize 'AND' because of operator precedence
					ans = dp(i, j+2) || isRuneMatch && dp(i+1, j)
				} else {
					// if not *, and this rune is match, check next rune (i+1,j+1)
					ans = isRuneMatch && dp(i+1, j+1)
				}
			}

			// memoize answer
			memo[indicesKey] = ans
		}

		return memo[indicesKey]
	}

	return dp(0, 0)
}
