package palindromenumber

func isPalindrome(x int) bool {
	if x >= 0 {
		i := 1
		j := 1

		for j*10 <= x {
			j *= 10
		}

		for i < j {
			if (x/j)%10 != (x%(i*10))/i {
				return false
			} else {
				i *= 10
				j /= 10
			}
		}

		return true

	} else {
		return false
	}
}
