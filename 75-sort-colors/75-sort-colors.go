package sortcolors

func sortColors(nums []int) {
	var bot, mid, top = -1, 0, len(nums)
	for mid < top && nums[mid] != 1 {
		if nums[mid] == 0 {
			bot++
			mid++
		} else {
			top--
			nums[mid], nums[top] = nums[top], nums[mid]
		}
	}
	for mid < top {
		if nums[mid] == 0 {
			bot++
			nums[bot], nums[mid] = nums[mid], nums[bot]
		} else if nums[mid] == 2 {
			top--
			nums[top], nums[mid] = nums[mid], nums[top]
		} else {
			mid++
		}
	}
}

// func sortColors(nums []int) {
// 	var zeroCount, oneCount int = 0, 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			zeroCount++
// 		} else if nums[i] == 1 {
// 			oneCount++
// 		}
// 	}
// 	for i := 0; i < zeroCount; i++ {
// 		nums[i] = 0
// 	}
// 	for i := zeroCount; i < zeroCount+oneCount; i++ {
// 		nums[i] = 1
// 	}
// 	for i := zeroCount + oneCount; i < len(nums); i++ {
// 		nums[i] = 2
// 	}
// }
