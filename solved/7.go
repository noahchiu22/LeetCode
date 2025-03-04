package main

import (
	"math"
)

func reverse(x int) int {
	out := 0
	absNum := math.Abs(float64(x))
	logNum := int(math.Log10(absNum))

	for absNum > 0 {
		out += int(absNum) % 10 * int(math.Pow10(logNum))

		absNum /= 10
		logNum--
	}

	if x < 0 {
		out *= -1
	}

	if out > int(math.Pow(2, 31))-1 || out < -int(math.Pow(2, 31)) {
		out = 0
	}

	return out
}
