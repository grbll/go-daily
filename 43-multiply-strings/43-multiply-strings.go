package multiplystrings

func multiply(num1 string, num2 string) string {
	var n, m int = len(num1), len(num2)

	var out string
	var bytes1, bytes2, work []byte = make([]byte, n, n), make([]byte, m, m), make([]byte, n+m, n+m)
	var temp, carry byte = 0, 0

	work = make([]byte, n+m, n+m)
	for i, num := range []byte(num1) {
		bytes1[i] = num - 48
		work[i] = 0
	}
	for j, num := range []byte(num2) {
		bytes2[j] = num - 48
		work[n+j] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j-- {
			temp = bytes2[m-1-i]*bytes1[n-1-j] + work[n+m-1-i-j] + carry
			work[n+m-1-i-j] = temp % 10
			carry = temp / 10
		}
		//bytes2[i]*bytes1
	}

	out = string(work)
	return out
}
