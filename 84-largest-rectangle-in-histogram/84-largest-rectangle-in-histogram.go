package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	var seen []int = make([]int, 1, len(heights))
	var best, start, curr int = heights[0], 0, 0
	for ind := 1; ind < len(heights); ind++ {
		for len(seen) > 1 && heights[seen[len(seen)-1]] > heights[ind] {
			curr = seen[len(seen)-1]
			seen = seen[:len(seen)-1]
			best = max(best, heights[curr]*(start-seen[len(seen)-1]))
		}
		if heights[seen[len(seen)-1]] > heights[ind] {
			best = max(best, heights[seen[0]]*(start+1))
			seen = seen[:len(seen)-1]
		}
		seen = append(seen, ind)
		best = max(best, heights[ind])
		start = ind
	}
	return best
}

// func split(ranges *[][2]int, splitIndices []int) {
// 	var oldRanges [][2]int = make([][2]int, 0, len(*ranges))
// 	oldRanges = append(oldRanges, *ranges...)
// 	*ranges = (*ranges)[:0]
// 	for _, newSplit := range splitIndices {
// 		for newSplit >= oldRanges[0][1] {
// 			*ranges = append(*ranges, oldRanges[0])
// 			oldRanges = oldRanges[1:]
// 		}
// 		if newSplit > oldRanges[0][0] {
// 			*ranges = append(*ranges, [2]int{oldRanges[0][0], newSplit})
// 		}
// 		oldRanges[0][0] = newSplit + 1
// 		if oldRanges[0][0] == oldRanges[0][1] {
// 			oldRanges = oldRanges[1:]
// 		}
// 	}
// 	*ranges = append(*ranges, oldRanges...)
// }
//
// func largestRectangleArea(heights []int) int {
// 	var best int = 0
// 	var ranges [][2]int = make([][2]int, 0, len(heights))
// 	ranges = append(ranges, [2]int{0, len(heights)})
// 	var indexMap map[int][]int = make(map[int][]int)
// 	for ind, hei := range heights {
// 		if _, ex := indexMap[hei]; ex {
// 			indexMap[hei] = append(indexMap[hei], ind)
// 		} else {
// 			indexMap[hei] = []int{ind}
// 		}
// 	}
// 	slices.Sort(heights)
// 	for ind, hei := range heights {
// 		if ind == len(heights)-1 || hei != heights[ind+1] {
// 			for _, ran := range ranges {
// 				best = max(best, (ran[1]-ran[0])*hei)
// 			}
// 			split(&ranges, indexMap[hei])
// 		}
// 	}
// 	return best
// }

// func largestRectangleArea(heights []int) int {
// 	var best, count, end int = 0, 0, 1
// 	var segments [][2]int = make([][2]int, 0, len(heights)*len(heights))
// 	var seg [2]int = [2]int{0, len(heights)}
// 	segments = append(segments, seg)
//
// 	for len(segments) > 0 {
// 		count++
// 		end = len(segments)
// 		for i := 0; i < end; i++ {
// 			seg = segments[0]
// 			segments = segments[1:]
// 			for j := seg[0]; j <= seg[1]; j++ {
// 				if j == seg[1] || heights[j] == 0 {
// 					if j > seg[0] {
// 						segments = append(segments, [2]int{seg[0], j})
// 						if (j-seg[0])*count > best {
// 							best = (j - seg[0]) * count
// 						}
// 					}
// 					seg[0] = j + 1
// 				}
// 				if j < seg[1] {
// 					heights[j]--
// 				}
// 			}
// 		}
// 	}
// 	return best
// }
