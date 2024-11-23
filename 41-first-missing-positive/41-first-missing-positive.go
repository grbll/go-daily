package firstmissingpositive

func heapRise(nums []int, tail int) {
	var curr, next int = tail - 1, (tail - 2) / 2
	for curr > 0 && nums[curr] < nums[next] {
		nums[curr], nums[next] = nums[next], nums[curr]
		curr = next
		next = (curr - 1) / 2
	}
}

func smallerChild(nums []int, curr, tail int) (int, bool) {
	var next int = curr * 2
	if next > tail-2 {
		return 0, false
	} else if next == tail-2 {
		return next + 1, true
	} else {
		if nums[next+1] <= nums[next+2] {
			return next + 1, true
		} else {
			return next + 2, true
		}
	}

}
func heapSink(nums []int, tail int) {
	var curr int = 0
	for {
		if next, ok := smallerChild(nums, curr, tail); ok {
			if nums[curr] > nums[next] {
				nums[next], nums[curr] = nums[curr], nums[next]
				curr = next
			} else {
				break
			}
		} else {
			break
		}
	}
}

func firstMissingPositive(nums []int) int {
	n := len(nums)

	var look, tail, next int

	look = 1
	tail, next = 0, 0

	for next < n {
		if nums[next] <= 0 {
			next++
		} else {
			if nums[next] == look {
				look++
				next++
				for tail > 0 && nums[0] <= look {
					if nums[0] == look {
						look++
					}
					tail--
					nums[0] = nums[tail]
					heapSink(nums, tail)
				}
			} else {
				nums[tail] = nums[next]
				tail++
				heapRise(nums, tail)
				next++
			}
		}
	}
	return look
}
