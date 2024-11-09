package generateparantheses

import (
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		// {
		// 	name:     "n=0",
		// 	n:        0,
		// 	expected: []string{""},
		// },
		// {
		// 	name:     "n=1",
		// 	n:        1,
		// 	expected: []string{"()"},
		// },
		// {
		// 	name:     "n=2",
		// 	n:        2,
		// 	expected: []string{"(())", "()()"},
		// },
		{
			name:     "n=3",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "n=4",
			n:        4,
			expected: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)

			// Sort both slices to ensure we can compare regardless of order
			sort.Strings(result)
			sort.Strings(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("generateParenthesis(%d) = %v, want %v", tt.n, result, tt.expected)
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("generateParenthesis(%d) = %v, want %v", tt.n, result, tt.expected)
					break
				}
			}
		})
	}
}
