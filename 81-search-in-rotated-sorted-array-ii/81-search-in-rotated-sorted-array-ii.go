package searchinrotatedsortedarrayii

func search(nums []int, target int) bool {
	var candidates [][2]int = make([][2]int, 0, len(nums))
	candidates = append(candidates, [2]int{0, len(nums) - 1})
	var mid, bot, top int
	for len(candidates) > 0 {
		bot, top = candidates[len(candidates)-1][0], candidates[len(candidates)-1][1]
		candidates = candidates[:len(candidates)-1]
		for bot < top {
			mid = (bot + top) / 2
			if nums[bot] < nums[top] {
				if target > nums[mid] {
					bot = mid + 1
				} else {
					top = mid
				}
			} else if nums[bot] > nums[top] {
				if nums[bot] <= nums[mid] {
					if target >= nums[bot] && target <= nums[mid] {
						top = mid
					} else {
						bot = mid + 1
					}
				} else {
					if target <= nums[mid] || target >= nums[bot] {
						top = mid
					} else {
						bot = mid + 1
					}
				}
			} else {
				if nums[bot] < nums[mid] {
					if target <= nums[mid] && target >= nums[bot] {
						top = mid
					} else {
						bot = mid + 1
					}
				} else if nums[bot] > nums[mid] {
					if target <= nums[mid] || target >= nums[bot] {
						top = mid
					} else {
						bot = mid + 1
					}
				} else {
					candidates = append(candidates, [2]int{mid + 1, top})
					top = mid
				}
			}
		}
		if nums[bot] == target {
			return true
		}
	}
	return false
}
