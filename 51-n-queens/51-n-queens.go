package nqueens

func check(rowMatrix []int) bool {
	var n int = len(rowMatrix) - 1
	for row, col := range rowMatrix[:n] {
		if col == rowMatrix[n] || col == rowMatrix[n]-(n-row) || col == rowMatrix[n]+(n-row) {
			return false
		}
	}
	return true
}

func buildString(rowMatrix []int) []string {
	var n int = len(rowMatrix)
	var line []byte = make([]byte, 0, n)
	var stringMatrix []string = make([]string, 0, n)
	for _, col := range rowMatrix {
		for i := 0; i < n; i++ {
			if i == col {
				line = append(line, 'Q')
			} else {
				line = append(line, '.')
			}
		}
		stringMatrix = append(stringMatrix, string(line))
		line = line[:0]
	}
	return stringMatrix
}

func solveNQueens(n int) [][]string {
	var out [][]string = make([][]string, 0)

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
				out = append(out, buildString(rowMatrix))
				rowMatrix = rowMatrix[:last]
			} else {
				rowMatrix = append(rowMatrix, -1)
			}
		}
	}

	return out
}
