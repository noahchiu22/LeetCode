package main

func orangesRotting(grid [][]int) int {
	count := 0

	for {
		hasChanged := false
		hasFresh := false
		tempGrid := make([][]int, len(grid))
		for i, row := range grid {
			tempGrid[i] = make([]int, len(row))
			copy(tempGrid[i], row)
		}

		for i, row := range tempGrid {
			for j, cell := range row {
				if cell == 2 {
					if i > 0 && tempGrid[i-1][j] == 1 {
						grid[i-1][j] = 2
						hasChanged = true
					}
					if i < len(tempGrid)-1 && tempGrid[i+1][j] == 1 {
						grid[i+1][j] = 2
						hasChanged = true
					}
					if j > 0 && tempGrid[i][j-1] == 1 {
						grid[i][j-1] = 2
						hasChanged = true
					}
					if j < len(row)-1 && tempGrid[i][j+1] == 1 {
						grid[i][j+1] = 2
						hasChanged = true
					}
				}
				if cell == 1 {
					hasFresh = true
				}
			}
		}

		if !hasChanged {
			if hasFresh {
				count = -1
			}
			break
		}

		count++
	}

	return count
}
