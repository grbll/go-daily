package setmatrixzeroes

var zeroes [200]int

func setZeroes(matrix [][]int) {
	var n, m int = len(matrix), len(matrix[0])
	var rowBool, colBool []bool = make([]bool, n, n), make([]bool, m, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rowBool[i] = true
				colBool[j] = true
			}
		}
	}
	var rows []int = make([]int, 0, n)
	for index, row := range rowBool {
		if row {
			matrix[index] = zeroes[:m]
		} else {
			rows = append(rows, index)
		}
	}
	for index, col := range colBool {
		if col {
			for _, row := range rows {
				matrix[row][index] = 0
			}
		}
	}
}
