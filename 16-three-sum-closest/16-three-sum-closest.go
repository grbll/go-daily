package threesumclosest

import (
	"sort"
)

func setAbsMin(a *int, b *int) {
	if *a <= 0 {
		if *b <= 0 {
			if *a < *b {
				*a = *b
			}
		} else {
			if *a < (-*b) {
				*a = *b
			}
		}
	} else {
		if *b <= 0 {
			if (-*b) < *a {
				*a = *b
			}
		} else {
			if *b < *a {
				*a = *b
			}
		}
	}
	return
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	lenN := len(nums)

	var index1, index2, index3 int
	var best, offset int

	best = nums[0] + nums[1] + nums[2] - target
	index1 = 0
	for index1 < lenN-2 {
		index2 = index1 + 1
		index3 = lenN - 1
		for index2 < index3 {
			offset = nums[index1] + nums[index2] + nums[index3] - target
			setAbsMin(&best, &offset)
			if offset <= 0 {
				index2++
			} else {
				index3--
			}
		}
		index1++
	}
	return target + best
}
