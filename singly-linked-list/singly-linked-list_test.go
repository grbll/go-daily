package singlylinkedlist

import (
	"testing"
)

func TestBuildListNode(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output *ListNode
	}{
		{
			name:   "Single element",
			input:  []int{1},
			output: &ListNode{Val: 1},
		},
		{
			name:   "Multiple elements",
			input:  []int{1, 2, 3, 4, 5},
			output: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
		{
			name:   "Empty list",
			input:  []int{},
			output: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildListNode(tt.input)
			if got, want := BuildSlice(result), tt.input; !slicesEqual(got, want) {
				t.Errorf("BuildListNode(%v) = %v, want %v", tt.input, got, want)
			}
		})
	}
}

func TestBuildSlice(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		output []int
	}{
		{
			name:   "Single element",
			input:  &ListNode{Val: 1},
			output: []int{1},
		},
		{
			name:   "Multiple elements",
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Empty list",
			input:  nil,
			output: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildSlice(tt.input)
			if got, want := result, tt.output; !slicesEqual(got, want) {
				t.Errorf("BuildSlice(%v) = %v, want %v", tt.input, got, want)
			}
		})
	}
}

func TestListNodeString(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		output string
	}{
		{
			name:   "Single element",
			input:  &ListNode{Val: 1},
			output: "[1]",
		},
		{
			name:   "Multiple elements",
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			output: "[1 2 3 4 5]",
		},
		{
			name:   "Empty list",
			input:  nil,
			output: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.output {
				t.Errorf("String() = %v, want %v", result, tt.output)
			}
		})
	}
}

// Helper function to check equality between slices of integers
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
