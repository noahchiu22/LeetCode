package main

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Println("s:", s, "p:", p)
	type runeProps struct {
		value      byte
		isFlexible bool
		count      int
	}

	pattern := []runeProps{}

	// generate pattern
	i := 0
	for i < len(p) {
		tempProps := runeProps{}
		// if next rune is *
		if i+1 < len(p) && p[i+1] == '*' {
			tempProps.isFlexible = true
		}

		// check duplicate rune
		if len(pattern) > 0 &&
			p[i] == pattern[len(pattern)-1].value &&
			tempProps.isFlexible == pattern[len(pattern)-1].isFlexible {
			pattern[len(pattern)-1].count++
		} else {
			// new rune
			tempProps.value = p[i]
			tempProps.count = 1

			pattern = append(pattern, tempProps)
		}

		i++

		// skip *
		if tempProps.isFlexible {
			i++
		}
	}

	for _, item := range pattern {
		fmt.Println("value:", string(item.value), "notNeedToCount:", item.isFlexible, "count:", item.count)
	}

	// check if s is valid
	i = 0
	j := 0 // for pattern
	for i < len(s) && j < len(pattern) {
		fmt.Println(string(s[i]), string(pattern[j].value))
		isRuneMatch := pattern[j].value == '.' || pattern[j].value == s[i]

		// if rune with * and rune is match, move s, else move p
		if pattern[j].isFlexible {
			if isRuneMatch {
				fmt.Println("move s")
				i++
			} else {
				fmt.Println("move p")
				j++
			}
			continue
		}

		dupEnd := i + pattern[j].count
		if dupEnd > len(s) {
			return false
		}

		// get fixed rune
		dupStr := ""
		if pattern[j].value == '.' {
			dupStr = s[i:dupEnd]
		} else {
			for idx := 0; idx < pattern[j].count; idx++ {
				dupStr += string(pattern[j].value)
			}
		}

		fmt.Println("s[i:dupEnd]", s[i:dupEnd], "dupStr", dupStr)

		// fixed rune check
		if s[i:dupEnd] != dupStr {
			return false
		}

		fmt.Println("move s+p")
		i += pattern[j].count
		j++
	}

	return i < len(s)
}
