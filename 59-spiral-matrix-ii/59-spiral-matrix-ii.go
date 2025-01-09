package spiralmatrixii

func generateMatrix(n int) [][]int {
	var count int = 0
	var outmatrix [][]int = make([][]int, n, n)
	for i := range outmatrix {
		outmatrix[i] = make([]int, n, n)
	}
	if n%2 == 1 {
		outmatrix[n/2][n/2] = n * n
	}
	for j := 0; j < n/2; j++ {
		for i := j; i < n-j-1; i++ {
			count++
			outmatrix[j][i] = count
		}
		for i := j; i < n-j-1; i++ {
			count++
			outmatrix[i][n-j-1] = count
		}
		for i := n - j - 1; i > j; i-- {
			count++
			outmatrix[n-j-1][i] = count
		}
		for i := n - j - 1; i > j; i-- {
			count++
			outmatrix[i][j] = count
		}
	}
	return outmatrix
}
