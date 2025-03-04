package main

func letterCombinations(digits string) (ans []string) {
	letterMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for i, c := range digits {
		tempAns := []string{}
		j := 0
		// run current ans or the first time
		for j < len(ans) || i == 0 {
			for _, letter := range letterMap[c] {
				if i == 0 {
					tempAns = append(tempAns, string(letter))
				} else {
					tempAns = append(tempAns, ans[j]+string(letter))
				}
			}
			// if it's the first time, break the loop
			if i == 0 {
				break
			}
			j++
		}
		ans = tempAns
	}

	return ans
}
