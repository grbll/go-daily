package climbingstairs

func nCr(n, k int) int {
	var out int = 1
	k = min(n-k, k)
	for i := 0; i < k; i++ {
		out *= (n - i)
		out /= (i + 1)
	}
	return out
}

func climbStairs(n int) int {
	var out int = 0
	for i := 0; i <= n/2; i++ {
		out += nCr(n/2+n%2+i, n%2+2*i)
	}
	return out
}
