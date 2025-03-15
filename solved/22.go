package main

func generateParenthesis(n int) []string {
	return dfs(n, "(")
}

func dfs(n int, bracket string) (out []string) {
	if len(bracket) == n*2-1 {
		return []string{bracket + ")"}
	}

	notPaired := 0
	headCount := 0
	for _, b := range bracket {
		if b == '(' {
			notPaired++
			headCount++
		} else {
			notPaired--
		}
	}
	if headCount < n {
		out = append(out, dfs(n, bracket+"(")...)
	}
	if notPaired > 0 {
		out = append(out, dfs(n, bracket+")")...)
	}

	return
}
