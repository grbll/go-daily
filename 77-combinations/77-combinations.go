package combinations

func combine(n int, k int) [][]int {
	var number int = 1
	for i := 0; i < k; i++ {
		number *= (n - i)
		number /= (i + 1)
	}
	var combination []int = make([]int, k, k)
	var out [][]int = make([][]int, 0, number)
	for i := 0; i < k; i++ {
		combination[i] = i + 1
	}
	out = append(out, append([]int{}, combination...))
	for len(out) < number {
		for i := k - 1; i >= 0; i-- {
			if combination[i] < n-(k-1-i) {
				combination[i]++
				for j := i + 1; j < k; j++ {
					combination[j] = combination[i] + (j - i)
				}
				out = append(out, append([]int{}, combination...))
				break
			}
		}
	}
	return out
}
