package permutationsii

import "slices"

func next(positions []int, end int) bool {
	for i := len(positions) - 1; i >= 0; i-- {
		if positions[i] < end {
			for j := len(positions) - 1; j >= i; j-- {
				positions[j] = positions[i] + 1
			}
			return true
		}
	}
	return false
}

func permuteUnique(nums []int) [][]int {
	var n, f int = len(nums), 1
	var numsMap map[int]int = make(map[int]int)
	var positions []int = make([]int, n, n)

	slices.Sort(nums)

	for index, num := range nums {
		f *= (index + 1)
		if index == 0 || num != nums[index-1] {
			numsMap[num] = 1
		} else {
			numsMap[num]++
			f /= numsMap[num]
		}
	}

	var out [][]int = make([][]int, 0, f)
	out = append(out, slices.Repeat([]int{nums[0]}, numsMap[nums[0]]))
	delete(numsMap, nums[0])
	var k, end int = 1, 1

	for num := range numsMap {
		positions = positions[:numsMap[num]]

		for next(positions, len(out[0])) {
			for i := 0; i < k; i++ {

				out = append(out, append(make([]int, 0, n), out[i][:positions[0]]...))
				for j := 0; j < len(positions)-1; j++ {
					out[end+i] = append(out[end+i], num)
					out[end+i] = append(out[end+i], out[i][positions[j]:positions[j+1]]...)
				}
				out[end+i] = append(out[end+i], num)
				out[end+i] = append(out[end+i], out[i][positions[len(positions)-1]:]...)

			}
			end += k
		}

		for i := 0; i < k; i++ {
			out[i] = append(slices.Repeat([]int{num}, numsMap[num]), out[i]...)
		}

		for index := range positions {
			positions[index] = 0
		}

		k = end
	}
	return out
}
