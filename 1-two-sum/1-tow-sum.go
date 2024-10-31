package twosum

func TwoSum(nums []int, target int) []int {
	hmap := make(map[int]int, len(nums))
	for index1, value := range nums {
		if index2, ok := hmap[value]; ok {
			return []int{index1, index2}
		} else {
			hmap[value-target] = index1
		}
	}
	return []int{}
}
