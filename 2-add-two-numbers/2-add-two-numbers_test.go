package addtwonumbers

import (
	. "github.com/grbll/go-daily/singly-linked-list"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Single node",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "Multiple nodes",
			input:    []int{3, 2, 1},
			expected: 3,
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := BuildListNode(tt.input)
			value := Get(&list)
			if value != tt.expected {
				t.Errorf("Get() = %d, want %d", value, tt.expected)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Same length lists without carry",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Different length lists",
			l1:       []int{9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 1},
		},
		{
			name:     "With carryover",
			l1:       []int{5, 6, 4},
			l2:       []int{5, 4, 6},
			expected: []int{0, 1, 1, 1},
		},
		{
			name:     "Empty first list",
			l1:       []int{},
			l2:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty second list",
			l1:       []int{1, 2, 3},
			l2:       []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both lists empty",
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := BuildListNode(tt.l1)
			l2 := BuildListNode(tt.l2)
			result := addTwoNumbers(l1, l2)
			got := BuildSlice(result)

			if !slicesEqual(got, tt.expected) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.expected)
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
