package zigzagconversion

func getIndex(row int, column int, mod int) *int {
	index := 2 * (column / mod) * mod

	if column%mod == 0 {
		index += row
	} else {
		if mod-row != column%mod {
			return nil
		}
		index = index + mod + column%mod
	}
	return &index
}

func convert(s string, numRows int) string {
	if numRows > 1 {
		mod := numRows - 1
		out := ""
		for row := 0; row <= mod; row++ {
			for column := 0; column <= (1+len(s)/(2*mod))*mod; column++ {
				index := getIndex(row, column, mod)
				if index != nil && *index < len(s) {
					out = out + string(s[*index])
				}
			}
		}
		return out
	} else {
		return s
	}
}
