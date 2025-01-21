package editdistance

func minDistance(word1 string, word2 string) int {
	var m, n int = len(word1) + 1, len(word2) + 1
	var editTable [][]int = make([][]int, n, n)
	for i := 0; i < n; i++ {
		editTable[i] = make([]int, m, m)
	}

	for i := 0; i < n; i++ {
		editTable[i][0] = i
	}
	for i := 0; i < m; i++ {
		editTable[0][i] = i
	}
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			if word1[j-1] == word2[i-1] {
				editTable[i][j] = editTable[i-1][j-1]
			} else {
				editTable[i][j] = myMin(editTable[i-1][j-1], editTable[i-1][j], editTable[i][j-1]) + 1
			}
		}

	}
	return editTable[n-1][m-1]
}

func myMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	}
	if b < c {
		return b
	}
	return c
}

// import slices
//
// func minDistance(word1 string, word2 string) int {
// 	if len(word2) < len(word1) {
// 		word1, word2 = word2, word1
// 	}
// 	if len(word1) == 0 {
// 		return len(word2)
// 	} else {
// 		var m, n = len(word1), len(word2)
// 		word2 = string(append(append(slices.Repeat([]byte{'#'}, m), []byte(word2)...), slices.Repeat([]byte{'#'}, m)...))
//
// 		var curr int
// 		var editPos []int = make([]int, n+2*m, n+2*m)
//
// 		for i := m; i < n+2*m; i++ {
// 			editPos[i] += i - (m - 1)
// 		}
//
// 		for i := 0; i < m; i++ {
// 			for j := n + 2*m - 1; j >= 0; j-- {
// 				curr = n + 2*m
// 				for l := 0; l < j; l++ {
// 					curr = min(curr, editPos[l]+(j-1-l))
// 				}
// 				if word2[j] == word1[i] {
// 					editPos[j] = min(curr, editPos[j]+1)
// 				} else {
// 					editPos[j] = min(curr, editPos[j]) + 1
// 				}
// 			}
// 		}
//
// 		for i := 0; i < m+n; i++ {
// 			editPos[i] += n - 1 + (m - i)
// 		}
//
// 		curr = n + 2*m
// 		for i := m; i < n+2*m-1; i++ {
// 			curr = min(curr, editPos[i])
// 		}
//
// 		return curr
// 	}
// }
