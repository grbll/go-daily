package permutations

func permute(nums []int) [][]int {
	var n, f int = len(nums), 1
	for i := 0; i < n; i++ {
		f *= (i + 1)
	}

	var out [][]int = make([][]int, f, f)

	for i := 0; i < f; i++ {
		out[i] = make([]int, 0, n)
	}

	out[0] = append(out[0], nums[0])
	f = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < f; k++ {
				out[(j+1)*f+k] = append(append(append(out[(j+1)*f+k], out[k][0:j]...), nums[i]), out[k][j:i]...)
			}
		}
		for k := 0; k < f; k++ {
			out[k] = append(out[k], nums[i])
		}
		f = f * (i + 1)
	}
	return out
}
