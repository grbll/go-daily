package minimumwindowsubstring

func minWindow(s string, t string) string {
	var tMap [173]int
	for _, v := range t {
		tMap[v]++
	}
	for i := 0; i < 120; i++ {
		if tMap[i] == 0 {
			tMap[i] = -1
		}
	}

	var index int = 0
	for index < len(s) && tMap[s[index]] < 0 {
		index++
	}
	var start, bestStart, bestLength, count int = index, index, len(s) + 1, len(t)
	for index < len(s) {
		tMap[s[index]]--
		if tMap[s[index]] < 0 && s[start] == s[index] {
			for {
				if tMap[s[start]] >= 0 {
					break
				} else {
					tMap[s[start]]++
					start++
				}
			}

		} else if tMap[s[index]] >= 0 {
			count--
		}
		if count == 0 && index-start < bestLength {
			bestLength = index - start + 1
			bestStart = start
		}
		index++
	}
	if bestLength < len(s)+1 {
		return s[bestStart : bestStart+bestLength]
	} else {
		return ""
	}
}

// func minWindow(s string, t string) string {
// 	var indexList []int = make([]int, 0, len(t)+1)
// 	var tMap map[byte]int = make(map[byte]int)
// 	var bot, top int = -1, len(s)
// 	for _, v := range []byte(t) {
// 		if _, e := tMap[v]; e {
// 			tMap[v]++
// 		} else {
// 			tMap[v] = 1
// 		}
// 	}
// 	for i, v := range []byte(s) {
// 		if num, e := tMap[v]; e {
// 			indexList = append(indexList, i)
// 			if num == 0 {
// 				for pos, index := range indexList {
// 					if s[index] == v {
// 						indexList = append(indexList[:pos], indexList[pos+1:]...)
// 						break
// 					}
// 				}
// 			} else {
// 				tMap[v]--
// 			}
// 			if len(indexList) == len(t) && indexList[len(t)-1]-indexList[0] < top-bot {
// 				bot, top = indexList[0], indexList[len(t)-1]
// 			}
// 		}
// 	}
// 	if bot >= 0 {
// 		return s[bot : top+1]
// 	} else {
// 		return ""
// 	}
// }

// type circle struct {
// 	values []int
// 	offset int
// }
//
// func NewCircle(length int) *circle {
// 	return &circle{values: make([]int, 0, length), offset: 0}
// }
//
// func (c *circle) add(index int) int {
// 	if len(c.values) < cap(c.values) {
// 		c.values = append(c.values, index)
// 		c.offset = (c.offset + 1) % cap(c.values)
// 		return -1
// 	} else {
// 		var temp int = c.values[c.offset]
// 		c.values[c.offset] = index
// 		c.offset = (c.offset + 1) % cap(c.values)
// 		return temp
// 	}
// }
//
// func (c *circle) get() (int, bool) {
// 	if len(c.values) == 0 {
// 		return 0, false
// 	} else if len(c.values) < cap(c.values) {
// 		return c.values[0], true
// 	} else {
// 		return c.values[c.offset], true
// 	}
// }
//
// type storage struct {
// 	letters map[byte]*circle
// 	lowest  int
// 	open    int
// }
//
// func (s *storage) add(b byte, index int) bool {
// 	if c, e := s.letters[b]; e {
// 		var low int = c.add(index)
// 		if low == s.lowest {
// 			s.lowest = index
// 			for _, c := range s.letters {
// 				if cLowest, cTest := c.get(); cTest {
// 					s.lowest = min(s.lowest, cLowest)
// 				}
// 			}
// 		} else if low < 0 {
// 			s.open--
// 		}
// 		if s.open == 0 {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func MakeStorage(tMap map[byte]int, start, number int, b byte) storage {
// 	var newStorage = storage{letters: make(map[byte]*circle), lowest: start, open: number - 1}
// 	for k, v := range tMap {
// 		newStorage.letters[k] = NewCircle(v)
// 	}
// 	newStorage.letters[b].add(start)
// 	return newStorage
// }
//
// func minWindow(s string, t string) string {
// 	var tMap map[byte]int = make(map[byte]int)
// 	var store storage
// 	for _, v := range []byte(t) {
// 		if _, e := tMap[v]; e {
// 			tMap[v]++
// 		} else {
// 			tMap[v] = 1
// 		}
// 	}
//
// 	var index = 0
// 	for index < len(s) {
// 		if _, e := tMap[s[index]]; e {
// 			store = MakeStorage(tMap, index, len(t), s[index])
// 			break
// 		}
// 		index++
// 	}
// 	if index == len(s) {
// 		return ""
// 	} else {
// 		var bot, top int = 0, len(s) - 1
// 		if store.open == 0 {
// 			return s[index : index+1]
// 		}
// 		index++
// 		for index < len(s) {
// 			if store.add(s[index], index) && index-store.lowest < top-bot {
// 				bot = store.lowest
// 				top = index
// 			}
// 			index++
// 		}
// 		if store.open == 0 {
// 			return s[bot : top+1]
// 		} else {
// 			return ""
// 		}
// 	}
// }

// func minWindow(s string, t string) string {
// 	var tMap [59]int
// 	var allMap [][59]int = make([][59]int, 2*len(s)-1, 2*len(s)-1)
//
// 	tMap[58] = len(t)
// 	for _, v := range []byte(t) {
// 		tMap[v-65]++
// 	}
// 	for i := 0; i < len(allMap); i++ {
// 		allMap[i] = tMap
// 	}
// 	for r := 0; r < len(s)+1; r++ {
// 		if r > 0 {
// 			for i := 0; i < 2*(len(s)-r)-1; i += 2 {
// 				if allMap[i+r][s[i/2]-65] > 0 {
// 					allMap[i+r][s[i/2]-65]--
// 					allMap[i+r][58]--
// 				}
// 			}
// 		}
// 		for i := 2 * r; i < 2*len(s)-1; i += 2 {
// 			if allMap[i-r][s[i/2]-65] > 0 {
// 				allMap[i-r][s[i/2]-65]--
// 				allMap[i-r][58]--
// 				if allMap[i-r][58] == 0 {
// 					return s[(i-2*r)/2 : i/2+1]
// 				}
// 			}
// 		}
// 	}
// 	return ""
// }
