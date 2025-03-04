package main

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0
	last := rune(s[0])
	for _, r := range s {
		num += romanMap[r]

		if (r == 'V' || r == 'X') && last == 'I' ||
			(r == 'L' || r == 'C') && last == 'X' ||
			(r == 'D' || r == 'M') && last == 'C' {
			num -= (romanMap[last] * 2)
		}
		last = r
	}

	return num
}
