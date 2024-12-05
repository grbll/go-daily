package groupanagrams

func groupAnagrams(strs []string) [][]string {
	var byteCounter [26]byte = [26]byte{}
	var anaMap map[[26]byte]int = make(map[[26]byte]int)
	var out [][]string = [][]string{}

	for _, ana := range strs {
		for _, letter := range []byte(ana) {
			byteCounter[letter-97]++
		}
		if pos, exists := anaMap[byteCounter]; exists {
			out[pos] = append(out[pos], ana)
		} else {
			anaMap[byteCounter] = len(out)
			out = append(out, []string{ana})
		}
		byteCounter = [26]byte{}
	}
	return out
}
