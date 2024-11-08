package lettercombinationofaphonenumber

import (
	"slices"
)

var numberToLetters map[byte][]byte = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var lenD = len(digits)
	if lenD != 0 {
		var out []string = []string{""}
		var currentLen int = 1

		for index1 := 0; index1 < lenD; index1++ {
			lenN := len(numberToLetters[digits[index1]])
			out = slices.Repeat(out, lenN)
			for index2 := 0; index2 < currentLen; index2++ {
				for index3 := 0; index3 < lenN; index3++ {
					out[index3*currentLen+index2] += string(numberToLetters[digits[index1]][index3])
				}
			}
			currentLen = len(out)
		}
		return out
	} else {
		return []string{}
	}
}
