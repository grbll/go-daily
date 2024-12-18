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

func solveNQueens(n int) [][]string {
	var out [][]string = make([][]string, 0)

	var stringMap map[int]string = make(map[int]string)
	for i := 0; i < n; i++ {
		var tempLine []byte = make([]byte, n, n)
		for j := 0; j < n; j++ {
			if i == j {
				tempLine[j] = 'Q'
			} else {
				tempLine[j] = '.'
			}
		}
		stringMap[i] = string(tempLine)
	}

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
				out = append(out, make([]string, n, n))
				for i := 0; i < n; i++ {
					out[len(out)-1][i] = stringMap[rowMatrix[i]]
				}
				rowMatrix = rowMatrix[:last]
			} else {
				rowMatrix = append(rowMatrix, -1)
			}
		}
	}

	return out
}
