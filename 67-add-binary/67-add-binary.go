package addbinary

// func addPowerOfTwo(a []byte,)
func addBinary(a string, b string) string {
	if len(b) < len(a) {
		a, b = b, a
	}
	var offset int = len(b) - len(a)

	var out []byte = append([]byte{'0'}, []byte(b)...)
	for index, digit := range a {
		if digit == '1' {
			for i := offset + index + 1; i >= 0; i-- {
				if out[i] == '0' {
					out[i] = '1'
					break
				} else {
					out[i] = '0'
				}
			}
		}
	}
	if out[0] == '0' {
		return string(out[1:])
	} else {
		return string(out)
	}
}
