package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	var n, m int = len(matrix), len(matrix[0])
	var out []int = make([]int, 0, n*m)
	var offset int = 0
	for offset*2+1 < n && offset*2+1 < m {
		for j := offset; j < m-offset; j++ {
			out = append(out, matrix[offset][j])
		}
		for j := offset + 1; j < n-offset; j++ {
			out = append(out, matrix[j][m-1-offset])
		}
		for j := m - 2 - offset; j > offset-1; j-- {
			out = append(out, matrix[n-1-offset][j])
		}
		for j := n - 2 - offset; j > offset; j-- {
			out = append(out, matrix[j][offset])
		}
		offset += 1
	}
	if (offset*2+1 == n && offset*2 < m) || (offset*2+1 == m && offset*2 < n) {
		for j := offset; j < m-offset; j++ {
			out = append(out, matrix[offset][j])
		}
		for j := offset + 1; j < n-offset; j++ {
			out = append(out, matrix[j][m-1-offset])
		}
	}
	return out
}
