package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	checkQueue, indexQueue := findFixedRune(p)

	fmt.Println(checkQueue)
	fmt.Println(indexQueue)
	fmt.Println("s: ", s)
	fmt.Println("p: ", p)
	sIndex, pIndex, cIndex := 0, 0, 0

	for sIndex < len(s) && pIndex < len(p) {
		fmt.Println("sIndex: ", sIndex, "pIndex: ", pIndex)

		if len(indexQueue) > 0 {
			// if is fixed pattern
			if indexQueue[cIndex] == pIndex {
				fmt.Println("fixed check")
				for i, r := range s[sIndex : sIndex+len(checkQueue[cIndex])] {
					checkRune := rune(checkQueue[cIndex][i])
					fmt.Println("r: ", string(r), "checkRune: ", string(checkRune))
					if r != checkRune && checkRune != '.' {
						return false
					}
				}
				sIndex += len(checkQueue[cIndex])
				pIndex += len(checkQueue[cIndex])
				cIndex++
				continue
			}

			// not fixed pattern
			if strings.Contains(checkQueue[cIndex], string(p[pIndex])) || p[pIndex] == '.' {
				// find duplicate s substring
				substring := ""
				checkRune := s[sIndex]
				for i := sIndex; i < len(s); i++ {
					if s[i] != checkRune {
						substring = s[sIndex:i]
						break
					}
					if i == len(s)-1 {
						substring = s[sIndex:]
					}
				}

				fmt.Println("substring: ", substring)

				// !!!!!!!!!!!!!!!!!!!!!!!!不行這樣檢查，有時候是要看substring完的下一筆rune
				// if len substring < len checkQueue[cIndex] return false
				if len(substring) < len(checkQueue[cIndex]) {
					return false
				}

				sIndex += len(substring)
				pIndex += (indexQueue[cIndex] + len(checkQueue[cIndex]))
				cIndex++
				continue
			}
		}

		// not fixed pattern

		// if match move sIndex
		if s[sIndex] == p[pIndex] || p[pIndex] == '.' {
			fmt.Println("s move")
			sIndex++
		} else {
			fmt.Println("p move")
			// if not match move pIndex
			// move 2 because of '*'
			pIndex += 2
		}
	}

	fmt.Println("sIndex: ", sIndex, "len(s): ", len(s))
	fmt.Println("cIndex: ", cIndex, "len(indexQueue): ", len(indexQueue))
	// if all check return true, not all check return false
	return sIndex >= len(s) && cIndex >= len(indexQueue)
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
