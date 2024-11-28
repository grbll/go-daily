package wildcardmatching

func merge(ranges, cutter, newRanges []int) []int {
	newRanges = []int{}
	var indexRanges, indexCutter = 0, 0
	var upper, lower *[]int
	var indexLower, indexUpper *int

	for indexRanges < len(ranges) && indexCutter < len(cutter) {
		if ranges[indexRanges+1]-1 < cutter[indexCutter] {
			indexRanges += 2
		} else if cutter[indexCutter+1]-1 < ranges[indexRanges] {
			indexCutter++
		} else {
			if ranges[indexRanges] < cutter[indexCutter] {
				indexLower = &indexCutter
				lower = &cutter
			} else {
				indexLower = &indexRanges
				lower = &ranges
			}
			if ranges[indexRanges+1] < cutter[indexRanges+1] {
				indexUpper = &indexRanges
				upper = &ranges
			} else {
				indexUpper = &indexCutter
				upper = &cutter
			}
			newRanges = append(newRanges, (*lower)[*indexLower]+1, (*upper)[(*indexUpper)+1]+1)
			*indexUpper += 2
		}
	}

	return newRanges
}

// func isMatch(s string, p string) bool { // new approach, marking s instead of p. old approachl below.
// 	var n int = len(s)
//
// 	var ranges, buildRanges []int = make([]int, 0, n+2), make([]int, 0, n+2)
// 	var cutterMap map[byte][]int
// 	var cutter []int
// 	var rangeIndex int
//
// 	ranges = []int{0, 1}
// 	for _, symbol := range []byte(p) {
// 		if symbol == '*' {
// 			ranges = []int{ranges[0], n + 1}
// 		} else {
//
// 			cutter, exists := cutterMap[symbol]
// 			if !exists {
// 				cutterMap[symbol] = make([]int, 0, n)
// 			}
//
// 			buildRanges = []int{}
// 			rangeIndex = 0
// 			for cutterIndex, boarder := range cutter {
// 				for boarder <
// 			}
// 		}
// 		if len(ranges) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// var mustConsume, canConsume int = 0, 0
// for _, symbol := range []byte(p) {
// 	if symbol == '*' {
// 		canConsume = n
// 	} else {
// 		mustConsume++
// 		canConsume = min(canConsume+1, n)
// 		if symbol != '?' {
// 			for canConsume >= mustConsume {
// 				if s[canConsume-1] == symbol {
// 					break
// 				} else {
// 					canConsume--
// 				}
// 			}
// 			for mustConsume <= canConsume {
// 				if s[mustConsume-1] == symbol {
// 					break
// 				} else {
// 					mustConsume++
// 				}
// 			}
// 		}
// 	}
// 	if mustConsume > canConsume {
// 		return false
// 	}
// }

// func isMatch(s string, p string) bool {
// 	var n = len(p)
//
// 	var positions []bool
//
// 	positions = make([]bool, n+1, n+1)
// 	positions[0] = true
//
// 	for index := 0; index < n; index++ {
// 		if p[index] != '*' {
// 			break
// 		} else {
// 			positions[index+1] = true
// 		}
// 	}
//
// 	for _, letter := range []byte(s) {
// 		positions[n] = false
// 		for index := n; index > 0; index-- {
// 			if positions[index-1] {
// 				switch p[index-1] {
// 				case '*':
// 					continue
// 				case '?':
// 					positions[index-1] = false
// 					positions[index] = true
// 				case letter:
// 					positions[index-1] = false
// 					positions[index] = true
// 				default:
// 					positions[index-1] = false
// 				}
// 			}
// 		}
// 		for index := 0; index < n; index++ {
// 			if positions[index] && p[index] == '*' {
// 				positions[index+1] = true
// 			}
// 		}
// 	}
//
// 	return positions[n]
// }
