package twosum_test

import (
	twosum "github.com/grbll/go-daily/001-two-sum"
	"reflect"
	"testing"
)

func TestLists(t *testing.T) {
	list1 := []int{2, 7, 11, 15}
	target1 := 9
	answer1 := []int{0, 1}

	test1 := twosum.TwoSum(list1, target1)
	if !reflect.DeepEqual(test1, answer1) {
		t.Errorf("Expected %v, got %v", answer1, test1)
	}
}
