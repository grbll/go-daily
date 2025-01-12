package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int = len(obstacleGrid), len(obstacleGrid[0])
	var pathGrid [][]int = make([][]int, m, m)
	for i := 0; i < m; i++ {
		pathGrid[i] = make([]int, n, n)
	}
	if obstacleGrid[m-1][n-1] == 0 {
		pathGrid[m-1][n-1] = 1
	}

	for i := n - 2; i > -m; i-- {
		for j := max(0, -i); j < min(m, n-i); j++ {
			if obstacleGrid[m-1-j][i+j] == 0 {
				if j > 0 {
					pathGrid[m-1-j][i+j] += pathGrid[m-1-j+1][i+j]
				}
				if i+j < n-1 {
					pathGrid[m-1-j][i+j] += pathGrid[m-1-j][i+j+1]
				}
			}
		}
	}
	return pathGrid[0][0]
}
