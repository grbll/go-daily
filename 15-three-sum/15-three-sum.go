package threesum

import "sort"

func jumpUp(nums []int, lenN *int, index *int) {
	for *index < *lenN-1 && nums[*index] == nums[*index+1] {
		*index++
	}
	*index++
}

func jumpDown(nums []int, index *int) {
	for *index > 0 && nums[*index] == nums[*index-1] {
		*index--
	}
	*index--
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	lenN := len(nums)
	out := make([][]int, 0)

	index1 := 0
	index2 := 0
	index3 := 0
	sum := 0

	for index1 < lenN-2 && (nums[index1] <= 0) {

		index2 = index1 + 1
		index3 = lenN - 1

		for index2 < index3 && nums[index3] >= 0 {

			sum = nums[index1] + nums[index2] + nums[index3]

			if sum < 0 {
				jumpUp(nums, &lenN, &index2)
			} else if sum > 0 {
				jumpDown(nums, &index3)
			} else {
				out = append(out, []int{nums[index1], nums[index2], nums[index3]})
				jumpUp(nums, &lenN, &index2)
				jumpDown(nums, &index3)
			}
		}
		jumpUp(nums, &lenN, &index1)
	}
	return out
}
