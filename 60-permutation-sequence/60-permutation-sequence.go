package permutationsequence

func getPermutation(n int, k int) string {
	k = k - 1
	var canditates, out []byte = make([]byte, n, n), make([]byte, 0, n)
	for i := 0; i < n; i++ {
		canditates[i] = byte(i + 49)
	}

	var curr, num int = 1, 0
	for i := 1; i < n; i++ {
		curr *= i
	}

	for i := n - 1; i > 0; i-- {
		num = k / curr
		out = append(out, canditates[num])
		canditates = append(canditates[:num], canditates[num+1:]...)
		k %= curr
		curr = curr / i
	}
	out = append(out, canditates[0])
	return string(out)
}
