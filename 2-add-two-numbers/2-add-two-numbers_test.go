package addtwonumbers

import (
	. "github.com/grbll/go-daily/singly-linked-list"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	value1 := addTwoNumbers(BuildListNode([]int{2, 1, 1}), BuildListNode([]int{1, 1, 1}))
	if value1.String() != "{3,2,2,}" {
		t.Errorf("want %v, got %v", BuildListNode([]int{3, 2, 2}).String(), value1.String())
	}
	value2 := addTwoNumbers(BuildListNode([]int{9, 9, 9}), BuildListNode([]int{1, 1, 1}))
	if value2.String() != "{0,1,1,1,}" {
		t.Errorf("want {0,1,1,1,}, got %v", value2.String())
	}
}
