package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	n := len(nums)
	offset := 0
	if n > 1 {
		for index, num := range nums[1:] {
			if nums[index-offset] == num {
				offset++
			} else {
				nums[index-offset+1] = num
			}
		}
	}
	return n - offset
}
