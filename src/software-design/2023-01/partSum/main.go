package main

import "fmt"

func main() {
	numOption := []int{3, 5, 7, 15, 20, 14, 23}
	limit := 56
	fmt.Println(findMaxDp(numOption, limit))
}

func findMaxDp(numOption []int, limit int) int {
	dpTable := initTable(len(numOption), limit + 1)
	rows := len(dpTable)
	cols := len(dpTable[0])

	for c := 0; c < cols; c++ {
		if numOption[0] > c {
			continue
		} else {
			dpTable[0][c] = numOption[0]
		}
	}

	for r := 1; r < rows; r++ {
		for c := 0; c < cols; c++ {
			candidate := numOption[r]
			notChoiceCaseMax := dpTable[r-1][c]
			choiceCaseMax := 0
			if candidate > c {
				continue
			} else {
				choiceCaseMax = dpTable[r-1][c-candidate] + candidate
			}
			dpTable[r][c] = max(notChoiceCaseMax, choiceCaseMax)
		}
	}

	return dpTable[len(numOption)-1][limit];
}

func initTable(row, col int) [][]int {
	table := make([][]int, row)
	for i := 0; i < row; i++ {
		table[i] = make([]int, col)
	}
	return table
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
