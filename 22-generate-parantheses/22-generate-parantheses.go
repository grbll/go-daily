package generateparantheses

func generateParenthesis(n int) []string {
	var paran [][]string = make([][]string, n+1, n+1)
	var nums []int = make([]int, n+1, n+1)
	paran[0], paran[1] = []string{""}, []string{"()"}
	nums[0], nums[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			nums[i] += nums[j] * nums[i-1-j]
		}
		paran[i] = make([]string, nums[i], nums[i])
		index := 0
		for j := 0; j < i; j++ {
			for _, sub1 := range paran[j] {
				for _, sub2 := range paran[i-1-j] {
					paran[i][index] = sub1 + "(" + sub2 + ")"
					index++
				}
			}
		}
	}
	return paran[n]
}
