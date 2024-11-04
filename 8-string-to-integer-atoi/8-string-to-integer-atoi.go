package stringtointegeratoi

const maxInt int64 = int64(^(^int32(0) << 31))
const minInt int64 = int64(^int32(0) << 31)

func myAtoi(s string) int {
	var out int64
	n := len(s)
	i := 0
	positive := true

	for i < n {
		if s[i] == 32 {
			i++
		} else {
			break
		}
	}

	if i < n && s[i] == 43 {
		i++
	} else if i < n && s[i] == 45 {
		positive = false
		i++
	}

	for i < n {
		if 48 <= s[i] && s[i] <= 57 {
			out = out*10 + int64(s[i]-48)
			if out > maxInt {
				if positive {
					return int(maxInt)
				} else {
					return int(minInt)
				}
			}
			i++
		} else {
			break
		}
	}
	if positive {
		return int(out)
	} else {
		return int(-out)
	}
}
