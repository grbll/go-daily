package multiplystrings

import "bytes"

func multiply(num1 string, num2 string) string {
	var n, m int = len(num1), len(num2)

	var out *bytes.Buffer = bytes.NewBuffer([]byte{})

	var bytes1, bytes2, work []byte = make([]byte, n, n), make([]byte, m, m), make([]byte, n+m, n+m)
	var temp, carry byte

	work = make([]byte, n+m, n+m)

	for i, num := range []byte(num1) {
		bytes1[n-1-i] = num - 48
		work[i] = 0
	}
	for j, num := range []byte(num2) {
		bytes2[m-1-j] = num - 48
		work[n+j] = 0
	}

	for i := 0; i < m; i++ {
		carry = 0
		for j := 0; j < n; j++ {
			temp = bytes2[i]*bytes1[j] + work[i+j] + carry
			work[i+j] = temp % 10
			carry = temp / 10
		}
		work[i+n] = carry
	}

	var firstNonZero = n + m - 1

	for firstNonZero > 0 {
		if work[firstNonZero] != 0 {
			break
		}
		firstNonZero--
	}

	for i := firstNonZero; i >= 0; i-- {
		out.WriteByte(work[i] + 48)
	}

	return out.String()
}
