package substringwithconcatenationofallwords

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected []int
	}{
		{
			name:     "three words, one occurances, mixed order",
			s:        "foobarfooloo",
			words:    []string{"loo", "bar", "foo"},
			expected: []int{3},
		},
		{
			name:     "three words, all duplicate, one occurances",
			s:        "foofoofoo",
			words:    []string{"foo", "foo", "foo"},
			expected: []int{0},
		},
		{
			name:     "three words, all duplicate, overlapping occurances",
			s:        "foofoofoofoofoo",
			words:    []string{"foo", "foo", "foo"},
			expected: []int{0, 3, 6},
		},
		{
			name:     "Test with overlapping words",
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "good"},
			expected: []int{8},
		},
		{
			name:     "Test with no matching substring",
			s:        "abcd",
			words:    []string{"efg", "hij"},
			expected: []int{},
		},
		{
			name:     "Test with empty string",
			s:        "",
			words:    []string{"foo", "bar"},
			expected: []int{},
		},
		{
			name:     "Test with words longer than string",
			s:        "a",
			words:    []string{"foo"},
			expected: []int{},
		},
		{
			name:     "error on submit 1",
			s:        "aaaaaaaaaaaaaa",
			words:    []string{"aa", "aa"},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.s, tt.words)
			sort.Ints(got)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.expected) {
				t.Errorf("%v: findSubstring(%v,%v) = %v, want %v", tt.name, tt.s, tt.words, got, tt.expected)
			}
		})
	}
}
