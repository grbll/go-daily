package longestvalidparantheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Empty string",
			s:    "",
			want: 0,
		},
		{
			name: "Single invalid parenthesis",
			s:    "(",
			want: 0,
		},
		{
			name: "Single valid pair",
			s:    "()",
			want: 2,
		},
		{
			name: "Mixed valid pairs",
			s:    "()()",
			want: 4,
		},
		{
			name: "Longest at the start",
			s:    "(()())()",
			want: 8,
		},
		{
			name: "Longest at the end",
			s:    "())(())",
			want: 4,
		},
		{
			name: "Nested parentheses",
			s:    "((()))",
			want: 6,
		},
		{
			name: "Invalid characters in between",
			s:    "(()))()",
			want: 4,
		},
		{
			name: "Alternating invalid and valid",
			s:    ")()())",
			want: 4,
		},
		{
			name: "All closed without opens",
			s:    "))))",
			want: 0,
		},
		{
			name: "All open without closes",
			s:    "((((",
			want: 0,
		},
		{
			name: "Complex pattern",
			s:    "(()(()))(()())",
			want: 14,
		},
		{
			name: "wrong pattern at end",
			s:    "()(()",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
