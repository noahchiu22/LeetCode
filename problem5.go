package main

func longestPalindrome(s string) string {
	lenS := len(s)

	maxStr := s[:1]
	for i := 0; i < lenS; i++ {
		var str string
		for l, r := i, i; l >= 0 && r < lenS && s[l] == s[r]; l, r = l-1, r+1 {
			str = s[l : r+1]
		}
		if len(maxStr) < len(str) {
			maxStr = str
		}

		for l, r := i, i+1; l >= 0 && r < lenS && s[l] == s[r]; l, r = l-1, r+1 {
			str = s[l : r+1]
		}
		if len(maxStr) < len(str) {
			maxStr = str
		}
	}
	return maxStr
}
