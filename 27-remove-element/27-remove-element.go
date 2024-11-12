package removeelement

func removeElement(nums []int, val int) int {
	n := len(nums)
	offset := 0
	index := 0
	for index < n-offset {
		if nums[index] == val {
			offset++
			nums[index] = nums[n-offset]
		} else {
			index++
		}
	}
	return n - offset
}
