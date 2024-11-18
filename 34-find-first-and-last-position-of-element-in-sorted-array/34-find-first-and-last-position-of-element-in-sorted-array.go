package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	n := len(nums)
	var index1, index2, pivot int = 0, n - 1, (n - 1) / 2
	for index2 > index1 && (nums[pivot] < target || nums[pivot] > target) {
		if nums[pivot] < target {
			index1 = pivot + 1
		} else {
			index2 = pivot - 1
		}
		pivot = index1 + (index2-index1)/2
	}
	if n == 0 || nums[pivot] != target {
		return []int{-1, -1}
	} else {
		var index3, index4 = pivot, pivot
		pivot = index1 + (index3-index1)/2
		for nums[index1] != target {
			if nums[pivot] < target {
				index1 = pivot + 1
			} else {
				index3 = pivot
			}
			pivot = index1 + (index3-index1)/2
		}
		pivot = index4 + (index2-index4+1)/2

		for nums[index2] != target {
			if nums[pivot] > target {
				index2 = pivot - 1
			} else {
				index4 = pivot
			}
			pivot = index4 + (index2-index4+1)/2
		}
		return []int{index1, index2}
	}
}
