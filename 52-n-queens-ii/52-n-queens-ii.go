package nqueensii

func check(rowMatrix []int) bool {
	var n int = len(rowMatrix) - 1
	for row, col := range rowMatrix[:n] {
		if col == rowMatrix[n] || col == rowMatrix[n]-(n-row) || col == rowMatrix[n]+(n-row) {
			return false
		}
	}
	return true
}

func totalNQueens(n int) int {
	var total = 0

	var last int
	var rowMatrix []int = make([]int, 0, n)
	rowMatrix = append(rowMatrix, -1)

	for len(rowMatrix) > 0 {
		last = len(rowMatrix) - 1
		rowMatrix[last]++
		if rowMatrix[last] == n {
			rowMatrix = rowMatrix[:last]
			continue
		}

		if check(rowMatrix) {
			if last == n-1 {
				total++
				rowMatrix = rowMatrix[:last]
			} else {
				rowMatrix = append(rowMatrix, -1)
			}
		}
	}

	return total
}
