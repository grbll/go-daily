package dividetwointegers

func divide(dividend int, divisor int) int {
	out := 0
	hold := 0
	negative := false
	position := int(1) << 30
	if dividend == -2147483648 {
		if divisor == 1 {
			return -2147483648
		}
		dividend = 0
		hold = 1
		negative = true
	} else {
		if dividend <= 0 {
			dividend = -dividend
			negative = true
		}
		for position != 0 && dividend&position == 0 {
			position >>= 1
		}
	}
	if divisor <= 0 {
		divisor = -divisor
		negative = !negative
	}

	for position != 0 {
		hold <<= 1
		out <<= 1
		if position&dividend != 0 {
			hold++
		}
		if hold >= divisor {
			hold -= divisor
			out++
		}
		position >>= 1
	}
	if negative {
		return -out
	} else {
		return out
	}
}
