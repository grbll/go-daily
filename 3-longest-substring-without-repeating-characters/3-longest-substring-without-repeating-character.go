package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	currentStart := 0
	maxSub := min(1, len(s))
	for index1 := 0; index1 < len(s); index1++ {
		for index2 := index1 - 1; index2 >= currentStart; index2-- { //this could be improved by using a hash map, however since this loop is bounded by the number of input symbols it would not improve the O-runtime
			if s[index1] == s[index2] {
				currentStart = index2 + 1
			}
		}
		maxSub = max(maxSub, index1-currentStart+1)
	}
	return maxSub
}
