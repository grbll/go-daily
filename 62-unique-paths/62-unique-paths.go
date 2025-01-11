package uniquepaths

func uniquePaths(m int, n int) int {
	m, n = min(m-1, n-1), max(m-1, n-1)
	var out int = 1
	for i := 0; i < m; i++ {
		out *= (n + (m - i))
		out /= (i + 1)
	}
	return out
}
