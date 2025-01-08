package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	var n = len(intervals)
	if n > 0 {
		var bot, mid1, mid2, top int = 0, n / 2, n / 2, n
		for bot < top {
			if newInterval[0] > intervals[mid1][0] {
				bot = mid1 + 1
			} else {
				top = mid1
			}
			mid1 = bot + (top-bot)/2
		}
		if mid1 > 0 && intervals[mid1-1][1] >= newInterval[0] {
			newInterval[0] = intervals[mid1-1][0]
			mid1--
		}
		bot, mid2, top = mid1, mid1+(n-mid1)/2, n
		for bot < top {
			if newInterval[1] < intervals[mid2][1] {
				top = mid2
			} else {
				bot = mid2 + 1
			}
			mid2 = bot + (top-bot)/2
		}
		if mid2 < n && intervals[mid2][0] <= newInterval[1] {
			newInterval[1] = intervals[mid2][1]
			mid2++
		}
		intervals = append(append(append([][]int{}, intervals[:mid1]...), newInterval), intervals[mid2:]...)
		return intervals
	} else {
		return [][]int{newInterval}
	}
}
