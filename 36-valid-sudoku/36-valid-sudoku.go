package validsudoku

func getSquare(row int, column int) int {
	return (row/3)*3 + (column / 3)
}

func isValidSudoku(board [][]byte) bool {
	var rowsMap, columnsMap, squaresMap []map[byte]struct{} = make([]map[byte]struct{}, 9), make([]map[byte]struct{}, 9), make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		rowsMap[i], columnsMap[i], squaresMap[i] = make(map[byte]struct{}), make(map[byte]struct{}), make(map[byte]struct{})
	}

	for i, row := range board {
		for j, num := range row {
			if num != '.' {
				_, existsRow := rowsMap[i][num]
				_, existsColumn := columnsMap[j][num]
				_, existsSquare := squaresMap[getSquare(i, j)][num]
				if existsRow || existsColumn || existsSquare {
					return false
				} else {
					rowsMap[i][num], columnsMap[j][num], squaresMap[getSquare(i, j)][num] = struct{}{}, struct{}{}, struct{}{}
				}
			}
		}
	}
	return true
}
