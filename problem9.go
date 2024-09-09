package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return palRecursion(x, int(math.Log10(float64(x))), 1)
}

func palRecursion(x, leftNum, rightNum int) bool {
	if leftNum < rightNum {
		return true
	}

	fmt.Println(x%int(math.Pow10(leftNum+1))/int(math.Pow10(leftNum)), x%int(math.Pow10(rightNum))/int(math.Pow10(rightNum-1)))
	if x%int(math.Pow10(leftNum+1))/int(math.Pow10(leftNum)) != x%int(math.Pow10(rightNum))/int(math.Pow10(rightNum-1)) {
		return false
	}

	return palRecursion(x, leftNum-1, rightNum+1)
}
