package removeduplicatesfromsortedarrayii

func rise(nums []int, index int) {
	var nextIndex int = (len(nums) - 1) - (len(nums)-2-index)/2
	for nums[index] < nums[nextIndex] {
		nums[index], nums[nextIndex] = nums[nextIndex], nums[index]
		index = nextIndex
		nextIndex = (len(nums) - 1) - (len(nums)-2-index)/2
	}
}

func sink(nums []int, end int) {
	var index, double int = len(nums) - 1, len(nums) - 1
	for {
		if double-2 >= end && nums[index] > nums[double-2] && nums[double-2] < nums[double-1] {
			nums[index], nums[double-2] = nums[double-2], nums[index]
			index = double - 2
		} else if double-1 >= end && nums[index] > nums[double-1] {
			nums[index], nums[double-1] = nums[double-1], nums[index]
			index = double - 1
		} else {
			break
		}
		double = len(nums) - 1 - (len(nums)-1-index)*2
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rise(nums, i)
	}
	nums[len(nums)-1], nums[0] = nums[0], nums[len(nums)-1]
	sink(nums, 1)
	nums[len(nums)-1], nums[1] = nums[1], nums[len(nums)-1]
	sink(nums, 2)
	var end, k int = 2, 2
	for end < len(nums) {
		end++
		nums[k], nums[len(nums)-1] = nums[len(nums)-1], nums[end-1]
		sink(nums, end)
		if nums[k] != nums[k-2] {
			k++
		}
	}
	return k
}
