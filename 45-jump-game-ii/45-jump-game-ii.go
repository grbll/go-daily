package jumpgameii

func jump(nums []int) int {
	var start, end, newEnd, count, n int = 0, 1, 0, 0, len(nums)
	for end < n {
		count++
		for start < end {
			newEnd = max(newEnd, start+nums[start]+1)
			start++
		}
		end = newEnd
	}
	return count
}
