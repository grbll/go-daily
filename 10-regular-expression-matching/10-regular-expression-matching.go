package regularexpressionmatching

func propagate(a *int32, b *int32, c *int32) int32 {
	if *a != 0 {
		var out int32 = *a
		var before int32 = int32(0)
		for before != out {
			before = out
			out = (((out & *b) << 1) | *a) & (*c - 1)
		}
		return out
	} else {
		return 0
	}

}

func buildOperators(p *string, lenPop *int32, operators *[27]int32) *[27]int32 {
	lenP := len(*p)
	index := 0

	for index < lenP-1 {
		if (*p)[index] == '.' {
			for j := 0; j < 26; j++ {
				operators[j] = operators[j] | *lenPop
			}
		} else {
			operators[(*p)[index]-'a'] = operators[(*p)[index]-'a'] | *lenPop
		}
		if (*p)[index+1] == '*' {
			operators[26] = operators[26] | *lenPop
			index++
		}
		*lenPop <<= 1
		index++
	}

	if (*p)[lenP-1] != '*' {
		if (*p)[index] == '.' {
			for j := 0; j < 26; j++ {
				operators[j] = operators[j] | *lenPop
			}
		} else {
			operators[(*p)[lenP-1]-'a'] = operators[(*p)[lenP-1]-'a'] | *lenPop
		}
		*lenPop <<= 1

	}
	return operators
}

func isMatch(s string, p string) bool {
	if p != "" {
		lenS := len(s)
		var lenPop int32 = int32(1)

		var allOperators [27]int32
		buildOperators(&p, &lenPop, &allOperators)

		var positions int32
		pPropagate := func() { positions = propagate(&positions, &allOperators[26], &lenPop) }

		lenPop <<= 1

		positions = int32(1)
		pPropagate()

		for i := 0; i < lenS; i++ {
			positions = (((positions & allOperators[s[i]-'a']) << 1) | (positions & allOperators[s[i]-'a'] & allOperators[26])) & (lenPop - 1)
			pPropagate()
		}

		return positions >= (lenPop >> 1)

	} else {
		return s == ""
	}
}
