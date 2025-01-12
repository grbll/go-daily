package minimumpathsum

func minPathSum(grid [][]int) int {
	var m, n int = len(grid), len(grid[0])
	for i := n - 2; i > -m; i-- {
		for j := max(0, -i); j < min(m, n-i); j++ {
			if j > 0 {
				if i+j < n-1 {
					grid[m-1-j][i+j] += min(grid[m-j][i+j], grid[m-j-1][i+j+1])
				} else {
					grid[m-1-j][i+j] += grid[m-j][i+j]
				}
			} else {
				grid[m-1-j][i+j] += grid[m-j-1][i+j+1]
			}
		}
	}
	return grid[0][0]
}
