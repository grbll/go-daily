package searchinrotatedsortedarray

func search(nums []int, target int) int {
	n := len(nums)
	if n > 0 {
		var index1, index2, pivot = 0, n, n / 2

		for index2-index1 > 2 && nums[index1] > nums[index2-1] {
			if nums[index1] > nums[pivot] {
				if target <= nums[pivot] || target >= nums[index1] {
					index2 = pivot + 1
				} else {
					index1 = pivot
				}
			} else {
				if target > nums[pivot] || target <= nums[index2-1] {
					index1 = pivot
				} else {
					index2 = pivot + 1
				}
			}
			pivot = index1 + (index2-index1)/2
		}
		for index2-index1 > 2 {
			if nums[pivot] <= target {
				index1 = pivot
			} else {
				index2 = pivot + 1
			}
			pivot = index1 + (index2-index1)/2
		}
		if nums[index1] == target {
			return index1
		} else if nums[index2-1] == target {
			return index2 - 1
		}
	}
	return -1
}
