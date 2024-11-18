package searchinsertposition

func searchInsert(nums []int, target int) int {
	var index1, index2, pivot int = 0, len(nums), len(nums) / 2
	for index2 > index1 {
		if nums[pivot] < target {
			index1 = pivot + 1
		} else {
			index2 = pivot
		}
		pivot = index1 + (index2-index1)/2
	}
	return index2
}
