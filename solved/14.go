package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstStr := strs[0]

	for i := range firstStr {
		for _, str := range strs {
			if i >= len(str) || str[i] != firstStr[i] {
				return firstStr[:i]
			}
		}
	}

	return firstStr
}
