package wildcardmatching

// import "slices"
//
// type store struct {
// 	jumps []int
// }
//
// func buildMapper(p string) *store {
// 	var n int = len(p)
// 	var pStore *store = &store{}
// 	var count int
//
// 	for i := 0; i < n; i++ {
// 		if p[i] != '*' {
// 			count++
// 		}
// 	}
//
// 	pStore.jumps = make([]int, count+1, count+1)
// 	pStore.jumps[count] = n
// 	count = 0
//
// 	for i := 0; i < n; i++ {
// 		if p[i] != '*' {
// 			pStore.jumps[count] = i
// 			count++
// 		}
// 	}
// 	return pStore
// }
//
// func (pStore *store) jump(pPositions []bool) {
// 	var n = len(pPositions)
// 	var index, next int
// 	index = 0
// 	for index < n {
// 		if pPositions[index] {
// 			next, _ = slices.BinarySearch(pStore.jumps, index)
// 			for index <= pStore.jumps[next] {
// 				pPositions[index] = true
// 				index++
// 			}
// 		} else {
// 			index++
// 		}
// 	}
// }

func isMatch(s string, p string) bool {
	var n = len(p)

	var positions []bool

	positions = make([]bool, n+1, n+1)
	positions[0] = true

	for index := 0; index < n; index++ {
		if p[index] != '*' {
			break
		} else {
			positions[index+1] = true
		}
	}

	for _, letter := range []byte(s) {
		positions[n] = false
		for index := n; index >= 0; index-- {
			if positions[index] {
				switch p[index] {
				case '*':
					continue
				case '?':
					positions[index] = false
					positions[index+1] = true
				case letter:
					positions[index] = false
					positions[index+1] = true
				default:
					positions[index] = false
				}
			}
		}
		for index := 0; index < n; index++ {
			if positions[index] && p[index] == '*' {
				positions[index+1] = true
			}
		}
	}

	return positions[n]
}
