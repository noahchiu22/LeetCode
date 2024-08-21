package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var tempS string = ""
	var max int = 0
	for i := range s {
		tempS = ""
		for j := i; j < len(s); j++ {
			if !strings.Contains(tempS, string(s[j])) {
				tempS += string(s[j])
			} else {
				break
			}
		}
		if len(tempS) > max {
			max = len(tempS)
		}
	}
	return max
}
