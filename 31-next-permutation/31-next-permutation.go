package nextpermutation

func nextPermutation(nums []int) {
	n := len(nums)
	index1 := n - 1
	if index1 >= 1 {
		for index1 > 0 && nums[index1-1] >= nums[index1] {
			index1--
		}
		for j := (n - index1) / 2; j > 0; j-- {
			nums[index1-1+j], nums[n-j] = nums[n-j], nums[index1-1+j]
		}
		if index1 > 0 {
			index2 := index1
			for index2 < n && nums[index1-1] >= nums[index2] {
				index2++
			}
			nums[index1-1], nums[index2] = nums[index2], nums[index1-1]
		}
	} else {
		return
	}
}
