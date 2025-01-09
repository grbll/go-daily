package lengthoflastword

func lengthOfLastWord(s string) int {
	var n int = len(s) - 1
	for n >= 0 && s[n] == ' ' {
		n--
	}
	var b int = n
	for n >= 0 && s[n] != ' ' {
		n--
	}
	return b - n
}
