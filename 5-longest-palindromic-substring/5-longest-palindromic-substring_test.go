package longestpalindromicsubstring

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaaa", "aaaa"},
		{"a", "a"},             // Single character
		{"ab", "a"},            // Two different characters
		{"ccc", "ccc"},         // Uneven substring better than even one
		{"cbbd", "bb"},         // Uneven substring better than even one
		{"aba", "aba"},         // Palindromic
		{"abba", "abba"},       // Even-length palindrome
		{"racecar", "racecar"}, // Longer palindrome
		{"abcdef", "a"},        // No palindrome longer than 1
		{"abacaba", "abacaba"}, // Palindrome with sub-palindromes
	}

	for _, tt := range tests {
		result := longestPalindrome(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("For input %q, expected %v, but got %v", tt.input, tt.expected, result)
		}
	}
}
