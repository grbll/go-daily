package foursum

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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	lenN := len(nums)

	var out [][]int = [][]int{}
	var i1, i2, i3, i4 int
	var sum int

	i1 = 0
	for i1 < lenN-2 {
		i2 = i1 + 1
		for i2 < lenN-1 {
			i3 = i2 + 1
			i4 = lenN - 1
			for i3 < i4 {
				sum = nums[i1] + nums[i2] + nums[i3] + nums[i4]
				if sum < target {
					jumpUp(nums, &lenN, &i3)
				} else if sum > target {
					jumpDown(nums, &i4)
				} else {
					out = append(out, []int{nums[i1], nums[i2], nums[i3], nums[i4]})
					jumpUp(nums, &lenN, &i3)
					jumpDown(nums, &i4)
				}
			}
			jumpUp(nums, &lenN, &i2)
		}
		jumpUp(nums, &lenN, &i1)

	}
	return out
}
