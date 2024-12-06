package powxn

func myPow(x float64, n int) float64 {
	var out float64 = 1

	if n < 0 { //flips n and x for division,
		x = 1 / x //works since 1/(x^n) == (1/x)^n
		n = -n
		if n == -2147483648 { //catch case where -n becomes 0
			out = x //since n is 0 in this case we will (s. loop) compute the square in each step of the loop. since the loop has 31 steps the result will be (1/x)^(2^31), which is (1/x)^n
		}
	}

	for i := 1 << 30; i > 0; i >>= 1 { //walk through the bits of the binary representation of n
		out *= out    //for each stem square current result (since out starts at 1 leading zeros will be ignored)
		if n&i == i { //if bit is a 1, we additionally need to multiply by x (base case n = 1)
			out *= x
		}
	}

	return out
}
