package validnumber

func isNumber(s string) bool {
	var sign, postsign, dot, predot, postdot, preexp, exp, postexp bool = false, true, false, false, true, false, false, true
	for _, digit := range s {
		if 47 < digit && digit < 58 {
			if dot {
				postdot = true
			} else {
				predot = true
			}
			if exp {
				postexp = true
			} else {
				preexp = true
			}
			postsign = true
			sign = true
		} else {
			if digit == '-' || digit == '+' {
				if sign {
					return false
				}
				postsign = false
				sign = true
			} else if digit == '.' {
				if dot || exp {
					return false
				} else {
					dot = true
					postdot = false
				}
				sign = true
			} else if digit == 'e' || digit == 'E' {
				if exp {
					return false
				}
				if dot && !(predot || postdot) {
					return false
				}
				exp = true
				postexp = false
				sign = false
			} else {
				return false
			}
		}
	}
	return postsign && (!exp || (preexp && postexp)) && (!dot || (predot || postdot))
}
