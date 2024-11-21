package combinationsum

import (
	"sort"
)

type indexToPebbles struct {
	index   int
	pebbles int
}

type indexStore struct {
	end   int
	store []indexToPebbles
}

func NewIndexStore(lenCandidates int) (*indexStore, bool) {
	if lenCandidates == 0 { //no solutions on empty input
		return nil, false
	} else {
		var newStore *indexStore = &indexStore{end: 0, store: make([]indexToPebbles, lenCandidates, lenCandidates)}
		newStore.store[0] = indexToPebbles{index: 0, pebbles: 1}
		return newStore, true
	}
}

func (is *indexStore) increase() { //add a pebbles at highest position
	is.store[is.end].pebbles++
}

func (is *indexStore) decrease() bool { //return value signals whether stor is empty after operation
	if is.store[is.end].index == len(is.store)-1 { //if the removed element is on last possible position
		is.end-- //delete position
	} else { //otherwise
		is.store[is.end].pebbles--         //remove one pebble
		if is.store[is.end].pebbles == 0 { //if position becomes empty
			is.end-- //delete it
		}
	}
	if is.end < 0 { //if no position remains
		return false //return false
	} else { //move one pebble one position up
		if is.store[is.end].pebbles == 1 { //if last pebble on position
			is.store[is.end].index++ //move position
		} else { //otherwise
			is.store[is.end].pebbles--                            //remove one pebble,
			is.end++                                              //create new postion
			is.store[is.end].index = is.store[is.end-1].index + 1 //one index after last postion
			is.store[is.end].pebbles = 1                          //and place one pebble there
		}
		return true
	}
}

func (is *indexStore) value(candidates []int) int {
	var value int = 0
	for i := 0; i <= is.end; i++ {
		value += candidates[is.store[i].index] * is.store[i].pebbles
	}
	return value
}

func (is *indexStore) calculateNumberOfPebbles() int {
	var number int = 0
	for i := 0; i < is.end+1; i++ {
		number += is.store[i].pebbles
	}
	return number
}

func (is *indexStore) build(candidates []int) []int {
	var numberOfPebbles int = is.calculateNumberOfPebbles()
	var out []int = make([]int, numberOfPebbles, numberOfPebbles)
	var runningIndex = 0
	for i := 0; i <= is.end; i++ {
		for j := 0; j < is.store[i].pebbles; j++ {
			out[runningIndex] = candidates[is.store[i].index]
			runningIndex++
		}
	}
	return out
}

func combinationSum(candidates []int, target int) [][]int {
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
			is.increase() //add a pebble
		} else {
			if value == target { //if target was hit
				out = append(out, is.build(candidates)) //add solution
			}
			nonempty = is.decrease() //remove one pebble an increas position of last, if no last remains set nonempty to false
		}
	}
	return out //return result
}
