package singlylinkedlist_test

import (
	"testing"

	sll "github.com/grbll/go-daily/singly-linked-list"
)

func TestBuildListNode(t *testing.T) {
	// Test case: Building a list with multiple values
	values := []int{1, 2, 3, 4}
	list := sll.BuildListNode(values)

	current := list
	for i, val := range values {
		if current == nil {
			t.Fatalf("Expected list to have at least %d elements, but found shorter list", len(values))
		}
		if current.Val != val {
			t.Errorf("Expected value at index %d to be %d, got %d", i, val, current.Val)
		}
		current = current.Next
	}

	if current != nil {
		t.Error("Expected end of list, but found extra nodes")
	}
}

func TestString(t *testing.T) {
	// Test case: Empty list
	emptyList := sll.BuildListNode([]int{})
	if emptyList.String() != "{}" {
		t.Errorf("Expected empty list string to be '{}', got %s", emptyList.String())
	}

	// Test case: List with elements
	values := []int{1, 2, 3, 4}
	list := sll.BuildListNode(values)
	expectedString := "{1,2,3,4,}"
	if list.String() != expectedString {
		t.Errorf("Expected list string to be %s, got %s", expectedString, list.String())
	}
}

func TestBuildListNode_EmptyInput(t *testing.T) {
	// Test case: Empty input
	list := sll.BuildListNode([]int{})
	if list != nil {
		t.Error("Expected nil list for empty input, got non-nil")
	}
}
