package main

import "math"

func intToRoman(num int) string {
	romanNum := []string{"I", "X", "C", "M"}
	romanNum5 := []string{"V", "L", "D"}

	romanStr := ""
	for i := len(romanNum) - 1; i >= 0; i-- {
		times := num / int(math.Pow10(i))
		num %= int(math.Pow10(i))
		tNum := ""
		for j := 0; j < times; j++ {
			if j == 3 {
				tNum = romanNum[i] + romanNum5[i]
				continue
			}
			if j == 4 {
				tNum = romanNum5[i]
				continue
			}
			if j == 8 {
				tNum = romanNum[i] + romanNum[i+1]
				continue
			}
			tNum += romanNum[i]
		}
		romanStr += tNum
	}

	return romanStr
}
