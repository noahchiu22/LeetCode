package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	checkQueue := []string{}
	i := len(p) - 1
	for i >= 0 {
		// skip rune with '*'
		if p[i] == '*' {
			i -= 2
			continue
		}

		for j := i; j >= 0; j-- {
			if j == 0 {
				checkQueue = append([]string{p[j : i+1]}, checkQueue...)
				i--
				break
			}
			if p[j] != p[j-1] {
				checkQueue = append([]string{p[j : i+1]}, checkQueue...)
				i -= (i - j + 1)
				break
			}
		}
	}

	for _, section := range checkQueue {
		fmt.Println(section)
	}

	fmt.Printf("s: %v\n", s)
	for i := 0; i < len(checkQueue)-1; i++ {
		fmt.Println("checkQueue[i]: ", checkQueue[i], "checkQueue[i+1]: ", checkQueue[i+1])

		index1 := indexOfNum(s, checkQueue[i], counts(checkQueue[:i], checkQueue[i])+1)
		index2 := indexOfNum(s, checkQueue[i+1], counts(checkQueue[:i+1], checkQueue[i+1])+1)

		fmt.Println("index1: ", index1, "index2: ", index2)
		if index1 == -1 || index2 == -1 {
			return false
		}
		if index1 > index2 {
			return false
		}
	}

	return true
}

// indexOfNum returns the index of the num-th occurrence of substr in s
func indexOfNum(s, substr string, num int) (index int) {
	index = strings.Index(s, substr)
	// end point
	if num == 1 {
		return
	}
	// if the length of s is less than the length of substr, return -1
	if len(s) < len(substr) {
		return -1
	}
	// recursive call
	nextIndex := indexOfNum(s[strings.Index(s, substr)+len(substr):], substr, num-1)
	// if the next index is -1, return -1
	if nextIndex == -1 {
		return -1
	}
	// add the length of substr and the next index
	index += (len(substr) + nextIndex)

	return
}

func counts[S []E, E comparable](s S, k E) int {
	h := map[E]int{}
	for _, v := range s {
		h[v] = h[v] + 1
	}
	return h[k]
}
