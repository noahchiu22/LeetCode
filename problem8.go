package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	numArr := []int{}
	out := 0.0
	isMinus := false
	isNum := false

	for _, r := range s {
		// ignore whitespace
		if r == 32 {
			if !isNum {
				continue
			} else {
				break
			}
		}
		// minus symbol
		if !isMinus && r == 45 {
			if isNum {
				break
			} else {
				isNum = true
				isMinus = true
				continue
			}
		}
		// plus symbol
		if r == 43 {
			if isNum {
				break
			} else {
				isNum = true
				continue
			}
		}
		if r >= 48 && r <= 57 {
			isNum = true
			if r == 48 && len(numArr) == 0 {
				continue
			}
			numArr = append(numArr, int(r-48))
		} else {
			break
		}
	}

	for i, num := range numArr {
		fmt.Println(math.Pow10(len(numArr) - (i + 1)))
		out += float64(num) * math.Pow10(len(numArr)-(i+1))

		if !isMinus && float64(out) > math.MaxInt32 {
			return math.MaxInt32
		}
		if isMinus && -float64(out) < -math.MaxInt32 {
			return -math.MaxInt32 - 1
		}
	}

	if isMinus {
		out *= -1
	}

	return int(out)
}
