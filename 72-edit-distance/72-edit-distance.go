package editdistance

import "slices"

func minDistance(word1 string, word2 string) int {
	if len(word2) < len(word1) {
		word1, word2 = word2, word1
	}
	if len(word1) == 0 {
		return len(word2)
	} else {
		var m, n = len(word1), len(word2)
		word2 = string(append(append(slices.Repeat([]byte{'#'}, m), []byte(word2)...), slices.Repeat([]byte{'#'}, m)...))

		var curr int
		var editPos []int = make([]int, n+2*m, n+2*m)

		for i := m; i < n+2*m; i++ {
			editPos[i] += i - (m - 1)
		}

		for i := 0; i < m; i++ {
			for j := n + 2*m - 1; j >= 0; j-- {
				curr = n + 2*m
				for l := 0; l < j; l++ {
					curr = min(curr, editPos[l]+(j-1-l))
				}
				if word2[j] == word1[i] {
					editPos[j] = min(curr, editPos[j]+1)
				} else {
					editPos[j] = min(curr, editPos[j]) + 1
				}
			}
		}

		for i := 0; i < m+n; i++ {
			editPos[i] += n - 1 + (m - i)
		}

		curr = n + 2*m
		for i := m; i < n+2*m-1; i++ {
			curr = min(curr, editPos[i])
		}

		return curr
	}
}
