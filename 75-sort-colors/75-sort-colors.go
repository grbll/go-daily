package sortcolors

func sortColors(nums []int) {
	var zeroCount, oneCount int = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else if nums[i] == 1 {
			oneCount++
		}
	}
	for i := 0; i < zeroCount; i++ {
		nums[i] = 0
	}
	for i := zeroCount; i < zeroCount+oneCount; i++ {
		nums[i] = 1
	}
	for i := zeroCount + oneCount; i < len(nums); i++ {
		nums[i] = 2
	}
}
