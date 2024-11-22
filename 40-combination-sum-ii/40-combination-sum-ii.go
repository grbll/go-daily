package combinationsumii

import "sort"

type indexStore struct {
	end   int
	store []int
}

func NewIndexStore(lenCandidates int) (*indexStore, bool) {
	if lenCandidates == 0 { //no solutions on empty input
		return nil, false
	} else {
		var newStore *indexStore = &indexStore{end: 0, store: make([]int, lenCandidates, lenCandidates)}
		newStore.store[0] = 0
		return newStore, true
	}
}

func (is *indexStore) increase(candidates []int) bool { //add a pebbles at highest position
	if is.store[is.end] == len(is.store)-1 {
		return is.decrease(candidates)
	} else {
		is.end++
		is.store[is.end] = is.store[is.end-1] + 1
		return true
	}
}

func (is *indexStore) decrease(candidates []int) bool { //return value signals whether stor is empty after operation
	if is.end == 0 {
		return false
	} else {
		is.end--
		for is.store[is.end] < len(candidates)-1 && candidates[is.store[is.end]] == candidates[is.store[is.end]+1] {
			is.store[is.end]++
		}
		is.store[is.end]++
		if is.store[is.end] == len(candidates) {
			return is.decrease(candidates)
		} else {
			return true
		}
	}
}

func (is *indexStore) value(candidates []int) int {
	var value int = 0
	for i := 0; i <= is.end; i++ {
		value += candidates[is.store[i]]
	}
	return value
}

func (is *indexStore) build(candidates []int) []int {
	var out []int = make([]int, is.end+1, is.end+1)
	for i := 0; i < is.end+1; i++ {
		out[i] = candidates[is.store[i]]
	}
	return out
}

func combinationSum2(candidates []int, target int) [][]int {
	//declare vars
	var out [][]int    //will store solutions
	var is *indexStore //contains the current placement of pebbles
	var nonempty bool  //tells us if finished or not
	var value int      //stores value during loop

	//initialize vars
	out = make([][]int, 0, 150) //at most 150 entries as per task

	is, nonempty = NewIndexStore(len(candidates))

	//sort candidates
	sort.Ints(candidates)

	//computation
	for nonempty { //as long as pebbles remain in store
		value = is.value(candidates) //set current value
		if value < target {          //if current value to small
			nonempty = is.increase(candidates) //add a pebble
		} else {
			if value == target { //if target was hit
				out = append(out, is.build(candidates)) //add solution
			}
			nonempty = is.decrease(candidates) //remove one pebble an increas position of last, if no last remains set nonempty to false
		}
	}
	return out //return result
}
