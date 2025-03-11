package main

func isValid(s string) bool {
	bracketMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := []string{}
	for _, char := range s {
		if len(stack) > 0 && string(char) == bracketMap[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(char))
		}
	}

	return len(stack) == 0
}
