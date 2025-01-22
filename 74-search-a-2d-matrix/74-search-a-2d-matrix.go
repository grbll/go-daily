package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	var bot, mid, top int = 0, len(matrix) / 2, len(matrix)
	for bot+1 < top {
		if matrix[mid][0] <= target {
			bot = mid
		} else {
			top = mid
		}
		mid = (bot + top) / 2
	}

	var row int = bot
	bot, mid, top = 0, len(matrix[row])/2, len(matrix[row])
	for bot+1 < top {
		if matrix[row][mid] <= target {
			bot = mid
		} else {
			top = mid
		}
		mid = (bot + top) / 2
	}
	return matrix[row][bot] == target
}
