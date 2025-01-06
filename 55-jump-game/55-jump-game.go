package jumpgame

func canJump(nums []int) bool {
	var n int = len(nums)
	var last int = -1
	for i := n - 2; i >= 0; i-- {
		if last < 0 && nums[i] == 0 {
			last = i
		} else {
			if i+nums[i] > last {
				last = -1
			}
		}
	}
	return last == -1
}
