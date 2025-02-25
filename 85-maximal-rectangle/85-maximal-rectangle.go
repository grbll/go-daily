package maximalrectangle

func sectionSize(mergeLine []int) int {
	var best int = 0
	for i := 0; i < len(mergeLine); i += 2 {
		best = max(best, mergeLine[i+1]-mergeLine[i])
	}
	return best
}

func sectionMerge(newLine []int, mergeLine *[]int) {
	if len(*mergeLine) != 0 {
		var indexNew, indexMerge int = 0, 0
		for indexNew < len(newLine) && indexMerge < len(*mergeLine) {
			if newLine[indexNew] >= (*mergeLine)[indexMerge+1] {
				*mergeLine = append((*mergeLine)[:indexMerge], (*mergeLine)[indexMerge+2:]...)
			} else if newLine[indexNew+1] <= (*mergeLine)[indexMerge] {
				indexNew += 2
			} else {
				(*mergeLine)[indexMerge] = max((*mergeLine)[indexMerge], newLine[indexNew])
				if newLine[indexNew+1] < (*mergeLine)[indexMerge+1] {
					*mergeLine = append((*mergeLine)[:indexMerge+2], append([]int{newLine[indexNew+1], (*mergeLine)[indexMerge+1]}, (*mergeLine)[indexMerge+2:]...)...)
					(*mergeLine)[indexMerge+1] = newLine[indexNew+1]
				}
				indexMerge += 2
			}
		}
		*mergeLine = (*mergeLine)[:indexMerge]
	}
}

func convertByteMatrixToRangeMatrix(matrix [][]byte) [][]int {
	var rangeMatrix [][]int = make([][]int, 0, len(matrix))
	var lookByte byte = '1'
	for _, byteRow := range matrix {
		var rangeRow []int = make([]int, 0, 2*((len(byteRow)+1)/2))
		lookByte = '1'
		for indCol, rowByte := range byteRow {
			if lookByte == rowByte {
				rangeRow = append(rangeRow, indCol)
				if lookByte == '1' {
					lookByte = '0'
				} else {
					lookByte = '1'
				}
			}
		}
		if len(rangeRow)%2 == 1 {
			rangeRow = append(rangeRow, len(byteRow))
		}
		rangeMatrix = append(rangeMatrix, rangeRow)
	}
	return rangeMatrix
}

func maximalRectangle(matrix [][]byte) int {
	var ind, height, best int = 0, 1, 0
	var rangeMatrix, mergeMatrix [][]int
	rangeMatrix = convertByteMatrixToRangeMatrix(matrix)
	var candidates []int = make([]int, 0, len(matrix))
	for i := 0; i < len(matrix); i++ {
		mergeMatrix = append(mergeMatrix, make([]int, 0, 2*((len(matrix[i])+1)/2)))
		mergeMatrix[i] = append(mergeMatrix[i], rangeMatrix[i]...)
		if len(mergeMatrix[i]) != 0 {
			candidates = append(candidates, i)
			best = max(best, sectionSize(mergeMatrix[i]))
		}
	}
	for len(candidates) > 0 {
		if candidates[len(candidates)-1]+height >= len(matrix) {
			candidates = candidates[:len(candidates)-1]
		}
		ind = 0
		for ind < len(candidates) {
			sectionMerge(rangeMatrix[candidates[ind]+height], &mergeMatrix[candidates[ind]])
			if len(mergeMatrix[candidates[ind]]) == 0 {
				candidates = append(candidates[:ind], candidates[ind+1:]...)
			} else {
				best = max(best, (height+1)*sectionSize(mergeMatrix[candidates[ind]]))
				ind++
			}
		}
		height++
	}
	return best
}
