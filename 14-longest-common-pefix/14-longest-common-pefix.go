package longestcommonpefix

func longestCommonPrefix(strs []string) string {
	lenS := len(strs)
	best := 0
	var index int

	for {
		index = 1
		if best >= len(strs[0]) {
			return strs[0][:best]
		} else {
			for index < lenS {
				if best >= len(strs[index]) || strs[index-1][best] != strs[index][best] {
					return strs[0][:best]
				}
				index++
			}
		}
		best++
	}
}
