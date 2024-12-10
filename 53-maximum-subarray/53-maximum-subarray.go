package maximumsubarray

func maxSubArray(nums []int) int {
	var index, temp, maxi int = 0, nums[0], nums[0]
	for index < len(nums)-1 {
		index++
		if temp+nums[index] < nums[index] {
			temp = nums[index]
		} else {
			temp += nums[index]
		}
		maxi = max(maxi, temp)
	}
	return maxi
}
