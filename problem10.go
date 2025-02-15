package main

import (
	"fmt"
)

type runeProps struct {
	value      string
	isFlexible bool
}

func isMatch(s string, p string) bool {
	fmt.Println("s:", s, "p:", p)

	pattern := []runeProps{}

	// generate pattern
	i := 0
	for i < len(p) {
		temp := runeProps{}
		// if next rune is *
		temp.isFlexible = i+1 < len(p) && p[i+1] == '*'

		//  if last and current both are not flexible, add to last value
		if len(pattern) > 0 && pattern[len(pattern)-1].isFlexible == temp.isFlexible && !temp.isFlexible {
			pattern[len(pattern)-1].value += string(p[i])
		} else {
			temp.value = string(p[i])

			pattern = append(pattern, temp)
		}

		i++

		// skip *
		if temp.isFlexible {
			i++
		}
	}

	a := match(s, pattern)
	fmt.Println("result:", a)

	return a == s
}

// let p try to match s
func match(s string, p []runeProps) (transP string) {
	fmt.Println("match start:", s, p)
	if len(s) == 0 {
		return ""
	}
	i := 0
	j := 0
	for i < len(s) && j < len(p) && len(transP) < len(s) {
		// fmt.Printf("i: %v, j: %v, s[i]: %v, p[j]: %v\n", i, j, string(s[i]), p[j].value)
		// if flexible
		if p[j].isFlexible {
			// find the next fixed value
			nextFixedIndex := 0
			for k := j; k < len(p); k++ {
				if !p[k].isFlexible {
					nextFixedIndex = k
					break
				}
			}

			index := windowSlide(i, s, p[nextFixedIndex].value)

			// if flexible turn into not flexible(after window slide)
			if index != -1 {
				result := match(s[index:], p[nextFixedIndex+1:])
				if s[:index]+result == s {
					transP += s[i:index] + result
					return
				}
			}

			// check if match move s, else move p
			if p[j].value == string(s[i]) || p[j].value == "." {
				transP += string(s[i])
				i++
			} else {
				j++
			}
		} else {
			// check if match move s and p, else return
			if checkTempMatch(i, s, p[j].value) {
				transP += string(s[i : i+len(p[j].value)])
				i += len(p[j].value)
				j++
			} else {
				return
			}
		}
	}

	fmt.Println(i, j)

	return
}

func windowSlide(start int, s, p string) int {
	// is not matchSwitch at the beginning
	matchSwitch := false
	i := start
	for p != "" && i+len(p) < len(s) {
		// false turn into true
		if !matchSwitch && checkTempMatch(i, s, p) {
			matchSwitch = true
		}
		// true turn into false
		if matchSwitch && !checkTempMatch(i, s, p) {
			return i + len(p) - 1
		}
		i++
	}

	return -1
}

func checkTempMatch(i int, s, p string) bool {
	// if not flexible
	isTempMatch := true
	// check if temp text match
	for k, r := range p {
		if i+k > len(s)-1 || (r != '.' && string(s[i+k]) != string(r)) {
			isTempMatch = false
			break
		}
	}

	return isTempMatch
}
