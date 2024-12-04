package rotateimage

func rotate(matrix [][]int) {
	var n int = len(matrix)
	for row := 0; row < n/2; row++ {
		for col := row; col < n-1-row; col++ {
			matrix[row][col], matrix[col][n-1-row], matrix[n-1-row][n-1-col], matrix[n-1-col][row] = matrix[n-1-col][row], matrix[row][col], matrix[col][n-1-row], matrix[n-1-row][n-1-col]
		}
	}
}
