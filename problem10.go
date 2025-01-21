package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	checkQueue, indexQueue := findFixedRune(p)

	fmt.Println(checkQueue)
	fmt.Println(indexQueue)
	fmt.Println("s: ", s)
	fmt.Println("p: ", p)
	sIndex, pIndex, cIndex := 0, 0, 0

	for sIndex < len(s) && pIndex < len(p) {
		// if is fixed pattern
		if indexQueue[cIndex] == pIndex {
			if s[sIndex] != p[pIndex] {
				return false
			}
			sIndex++
			pIndex++
			cIndex++
			continue
		}

		// not fixed pattern

		// if match move sIndex
		if s[sIndex] == p[pIndex] {
			sIndex++
		} else {
			// if not match move pIndex
			// move 2 because of '*'
			pIndex += 2
		}
	}

	// if not all check return false
	return sIndex < len(s) || cIndex < len(indexQueue)
}

// find fixed rune and index
func findFixedRune(p string) (checkQueue []string, indexQueue []int) {
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
				indexQueue = append([]int{j}, indexQueue...)
				i -= (i - j + 1)
				break
			}
			if p[j] != p[j-1] {
				checkQueue = append([]string{p[j : i+1]}, checkQueue...)
				indexQueue = append([]int{j}, indexQueue...)
				i -= (i - j + 1)
				break
			}
		}
	}

	return
}

// func gapCheck(s string, p string) bool {
// 	fmt.Println("gapCheck -----> ", "s: ", s, "p: ", p)
// 	if len(s) == 0 {
// 		return true
// 	}
// 	if len(p) == 0 {
// 		return false
// 	}

// 	pivot := 0
// 	t := '\n'
// 	for _, r := range s {
// 		if t != r {
// 			pivot = strings.Index(p[pivot:], string(r)) + 1
// 			if pivot == 0 {
// 				return false
// 			}
// 			t = r
// 		}
// 	}

// 	return true
// }
