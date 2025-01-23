package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isMatch("aaa", "a*a*aaa"))
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
	fmt.Println(isMatch("bbbba", ".*a*a"))
}
