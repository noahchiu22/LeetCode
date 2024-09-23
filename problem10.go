package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	fmt.Println(s)
	fmt.Println(p)
	pIndex, sIndex := len(p)-1, len(s)-1

	// comparing pattern
	for pIndex >= 0 {
		fmt.Println("sIndex:", sIndex, "pIndex:", pIndex)

		checkIndex := sIndex

		// if is flexible
		if p[pIndex] == '*' {
			for p[pIndex] == '*' {
				pIndex--
			}
			for p[pIndex] == '.' || s[sIndex] == p[pIndex] {
				sIndex--
				// if different, then break
				if sIndex == -1 {
					break
				}
			}
			pIndex--
		} else {
			if p[pIndex] != '.' && s[sIndex] != p[pIndex] {
				return false
			}
			pIndex--
			sIndex--
		}

		// if have compared all s, but p haven't
		if sIndex == -1 && pIndex != -1 {
			return isMatch(s[:checkIndex+1], p[:pIndex+1]+"*")
		}
	}

	return sIndex == -1
}

// func isMatch(s string, p string) bool {
// 	fmt.Println(s)
// 	fmt.Println(p)
// 	pIndex, sIndex, isMatch, isMultiple := 0, 0, true, false

// 	// comparing pattern
// 	for pIndex < len(p) {
// 		fmt.Println("sIndex:", sIndex, "pIndex:", pIndex)

// 		if pIndex == len(p)-1 {
// 			isMultiple = false
// 		} else {
// 			isMultiple = p[pIndex+1] == '*'
// 		}

// 		// check if have duplicate rune, * change with last duplicate rune
// 		if isMultiple && pIndex < len(p)-2 {
// 			lastIndex := pIndex + 2
// 			isChange := p[lastIndex] == p[pIndex]
// 			for lastIndex < len(p) {
// 				if p[lastIndex] != p[pIndex] {
// 					lastIndex--
// 					break
// 				}
// 				lastIndex++
// 			}
// 			if lastIndex == len(p) {
// 				lastIndex--
// 			}

// 			// if change index
// 			if isChange {
// 				p = p[:pIndex+1] + string(p[pIndex]) + p[pIndex+2:]
// 				p = p[:lastIndex] + "*" + p[lastIndex+1:]
// 				isMultiple = false
// 				fmt.Println(p)
// 			}
// 		}

// 		fmt.Println("isMultiple", isMultiple)

// 		if isMultiple {
// 			if sIndex > len(s)-1 {
// 				pIndex += 2
// 				continue
// 			}
// 			// if find a rune uncompared
// 			if p[pIndex] != '.' && s[sIndex] != p[pIndex] {
// 				pIndex += 2
// 			} else {
// 				sIndex++
// 			}
// 		} else {
// 			if sIndex > len(s)-1 {
// 				// offset
// 				if s[len(s)-1] == p[pIndex] && p[pIndex-1] == '*' {
// 					count := 0
// 					fixedCount := 0
// 					i := pIndex - 1
// 					for i >= 0 {
// 						if p[i] == '*' {
// 							i -= 2
// 						} else {
// 							if p[i] == p[pIndex] {
// 								fixedCount++
// 							} else {
// 								break
// 							}
// 							i--
// 						}
// 					}
// 					for pIndex < len(p) {
// 						if p[pIndex] != s[len(s)-1] {
// 							isMatch = false
// 							fmt.Println("trigger 1")
// 							break
// 						}

// 						count++

// 						fmt.Println("@@@@@@@@@@", count, len(s)-fixedCount)
// 						if count > len(s)-fixedCount {
// 							isMatch = false
// 							fmt.Println("trigger 2")
// 							return isMatch
// 						}

// 						pIndex++
// 					}
// 					continue
// 				} else {
// 					isMatch = false
// 					fmt.Println("trigger 3")
// 					break
// 				}
// 			}
// 			// if find a rune uncompared
// 			if p[pIndex] != '.' && s[sIndex] != p[pIndex] {
// 				isMatch = false
// 				fmt.Println("trigger 4")
// 				break
// 			}
// 			pIndex++
// 			sIndex++
// 		}
// 	}

// 	if pIndex < len(p) || sIndex < len(s) {
// 		isMatch = false
// 	}

// 	return isMatch
// }
