package twosum

func TwoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for index1 := 0; index1 < len(nums); index1++ {
		if index2, ok := hmap[nums[index1]]; ok {
			return []int{index2, index1}
		} else {
			hmap[target-nums[index1]] = index1
		}
	}
	return []int{}
}
