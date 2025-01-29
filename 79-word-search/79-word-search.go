package wordsearch

var shiftTable [4][2]int = [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

type entry struct {
	row   int
	col   int
	count int
}

func exist(board [][]byte, word string) bool {
	var m, n int = len(board), len(board[0])
	var newRow, newCol int = 0, -1
	var entries []entry = make([]entry, 0, len(word))
	var root entry = entry{row: newRow, col: newCol}

	for len(entries) > 0 || root.row*n+root.col < m*n-1 {
		if len(entries) == len(word) {
			return true
		}
		if len(entries) == 0 {
			newRow, newCol = root.row+(root.col+1)/n, (root.col+1)%n
			if board[newRow][newCol] == word[0] {
				entries = append(entries, entry{row: newRow, col: newCol, count: 4})
				board[newRow][newCol] = '#'
			}
			root.row = newRow
			root.col = newCol
		} else {
			var last *entry = &entries[len(entries)-1]
			for last.count >= 0 {
				last.count--
				if last.count == -1 {
					board[last.row][last.col] = word[len(entries)-1]
					entries = entries[:len(entries)-1]
				} else {
					newRow, newCol = last.row+shiftTable[last.count][0], last.col+shiftTable[last.count][1]
					if 0 <= newRow && newRow < m && 0 <= newCol && newCol < n && board[newRow][newCol] == word[len(entries)] {
						entries = append(entries, entry{row: newRow, col: newCol, count: 4})
						board[newRow][newCol] = '#'
						break
					}
				}
			}
		}
	}
	return false
}
