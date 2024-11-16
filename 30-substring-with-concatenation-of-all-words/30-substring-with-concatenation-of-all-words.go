package substringwithconcatenationofallwords

type counter struct {
	current  int
	capacity int
}

type modIndex struct{ val int }

func (a *modIndex) Down(n int) {
	a.val = (a.val + n - 1) % n
}

func findSubstring(s string, words []string) []int {
	n := len(words)
	m := len(words[0])
	k := len(s)
	var wordsMap map[string]*counter = make(map[string]*counter)
	var out []int = []int{}
	for _, word := range words {
		if currentCounter, ok := wordsMap[word]; ok {
			currentCounter.capacity++
		} else {
			wordsMap[word] = &counter{current: 0, capacity: 1}
		}
	}
	for j := 0; j < m; j++ {
		queue := make([]string, n, n)
		begQ := modIndex{val: 0}
		lenQ := 0
		for key := range wordsMap {
			wordsMap[key].current = 0
		}
		for i := k - m - j; i >= 0; i -= m {
			if currentCounter, ok := wordsMap[s[i:i+m]]; ok {
				for currentCounter.current >= currentCounter.capacity {
					lenQ--
					wordsMap[queue[(begQ.val+lenQ)%n]].current--
				}
				begQ.Down(n)
				lenQ++
				queue[begQ.val] = s[i : i+m]
				currentCounter.current++
				if lenQ == n {
					out = append(out, i)
				}
			} else {
				for lenQ > 0 {
					lenQ--
					wordsMap[queue[(begQ.val+lenQ)%n]].current--
				}
			}
		}
	}
	return out
}
