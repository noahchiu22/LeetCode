package main

func convert(s string, numRows int) string {
	letterRows := make([]string, numRows)

	for i, r := range s {
		// Decide the rowIndex of letter ((numsRows*2-1 = rowCount)-1 = rowIndex)
		rowIndex := 0
		if numRows > 1 {
			rowIndex = i % ((numRows*2 - 1) - 1)
		}
		if rowIndex < numRows {
			letterRows[rowIndex] += string(r)
		} else {
			// Letters in the middle of col need to be reversed
			letterRows[numRows-1-(rowIndex%numRows+1)] += string(r)
		}
	}

	output := ""
	for _, rowLetter := range letterRows {
		output += rowLetter
	}

	return output
}

// func convert(s string, numRows int) string {
// 	letterRows := make([]string, numRows)

// 	for i, r := range s {
// 		// Decide the rowIndex of letter ((numsRows*2-1 = rowCount)-1 = rowIndex)
// 		rowIndex := 0
// 		if numRows > 1 {
// 			rowIndex = i % ((numRows*2 - 1) - 1)
// 		}
// 		if rowIndex < numRows {
// 			letterRows[rowIndex] += string(r)
// 		} else {
// 			// Letters in the middle of col need to be reversed
// 			letterRows[numRows-1-(rowIndex%numRows+1)] += string(r)
// 		}
// 	}

// 	return strings.Join(letterRows, "")
// }
