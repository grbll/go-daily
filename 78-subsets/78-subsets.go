package subsets

func subsets(nums []int) [][]int {
	var out [][]int = make([][]int, 1, 1<<len(nums))
	var n int = 1
	for _, num := range nums {
		for i := 0; i < n; i++ {
			out = append(out, append(append([]int{}, out[i]...), num))
		}
		n = len(out)
	}
	return out
}
