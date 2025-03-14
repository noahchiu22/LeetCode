package main

import "slices"

func generateParenthesis(n int) []string {
	ans := []string{}
	bracket := ""
	for i := 0; i < n; i++ {
		bracket += "("
	}
	for i := 0; i < n; i++ {
		bracket += ")"
	}

	ans = append(ans, insidOut(bracket)...)

	bracket = ""
	for i := 0; i < n; i++ {
		bracket += "()"
	}

	if !slices.Contains(ans, bracket) {
		ans = append(ans, bracket)
	}

	return ans
}

func insidOut(in string) (out []string) {
	out = append(out, in)
	if in[1] == ')' || in[len(in)-2] == '(' {
		return
	}
	for i := 1; i < len(in)-1; i++ {
		if in[i] == '(' && in[i+1] == ')' {
			brackets := insidOut(in[:i] + ")(" + in[i+2:])
			for _, bracket := range brackets {
				if !slices.Contains(out, bracket) {
					out = append(out, bracket)
				}
			}
		}
	}
	return
}
