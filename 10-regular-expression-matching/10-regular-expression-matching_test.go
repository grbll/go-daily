package regularexpressionmatching

import (
	"testing"
)

func Test_propagate(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int32
		want    int32
	}{
		{
			name: "basic propagation with no cutoff",
			a:    0b00100000,
			b:    0b11110001,
			c:    0b10000000, // c as a power of two
			want: 0b01100000,
		},
		{
			name: "full propagation but limited by c",
			a:    0b00010000,
			b:    0b11111111,
			c:    0b01000000, // c as a power of two, restricts output
			want: 0b00110000, // output limited to values below 0b01000000
		},
		{
			name: "no propagation with zero in b",
			a:    0b00000001,
			b:    0b00000000,
			c:    0b10000000,
			want: 0b00000001, // no bits propagate due to zero in b
		},
		{
			name: "cutoff enforced early by c",
			a:    0b00100010,
			b:    0b11100111,
			c:    0b10000000, // c limits propagation to five positions
			want: 0b01101110, // propagation stops below 0b00100000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propagate(&tt.a, &tt.b, &tt.c)
			if got != tt.want {
				t.Errorf("propagate(%#08b, %#08b, %#08b) = %#08b, want %#08b", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}
func ptr(i int32) *int32 {
	return &i
}
func Test_buildOperators(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [27]int32
	}{
		{
			name:  "simple case with no stars",
			input: "abc",
			want:  [27]int32{1, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:  "case with a star",
			input: "ab*c",
			want:  [27]int32{1, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		},
		{
			name:  "multiple stars and characters",
			input: "a*b*c",
			want:  [27]int32{1, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3},
		},
		{
			name:  "ending with a character",
			input: "abc*",
			want:  [27]int32{1, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
		},
		{
			name:  "long input with multiple stars",
			input: "a*f*c*d*e*b*",
			want:  [27]int32{1, 32, 4, 8, 16, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var operators [27]int32
			p := &tt.input
			length := int32(1)
			got := buildOperators(p, &length, &operators)
			if *got != tt.want {
				t.Errorf("buildOperators(%q) = %v, want %v", tt.input, *got, tt.want)
			}
		})
	}
}
func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		// {
		// 	name: "exact match with no Kleene stars",
		// 	s:    "abc",
		// 	p:    "abc",
		// 	want: true,
		// },
		// {
		// 	name: "Kleene star matches zero occurrences",
		// 	s:    "abc",
		// 	p:    "ab*c",
		// 	want: true,
		// },
		// {
		// 	name: "Kleene star matches multiple occurrences",
		// 	s:    "abbbc",
		// 	p:    "ab*c",
		// 	want: true,
		// },
		// {
		// 	name: "Kleene star matches entire sequence",
		// 	s:    "aaaaa",
		// 	p:    "a*",
		// 	want: true,
		// },
		// {
		// 	name: "Kleene star with multiple patterns",
		// 	s:    "aabcc",
		// 	p:    "a*b*c*",
		// 	want: true,
		// },
		// {
		// 	name: "no match with additional characters",
		// 	s:    "abcd",
		// 	p:    "ab*e",
		// 	want: false,
		// },
		// {
		// 	name: "Kleene star pattern mismatch",
		// 	s:    "abcc",
		// 	p:    "ab*c",
		// 	want: false,
		// },
		// {
		// 	name: "Kleene star with non-repeating character",
		// 	s:    "ac",
		// 	p:    "a*b*c",
		// 	want: true,
		// },
		// {
		// 	name: "empty string with pattern containing only Kleene star",
		// 	s:    "",
		// 	p:    "a*",
		// 	want: true,
		// },
		{
			name: "empty string with empty pattern",
			s:    "",
			p:    "",
			want: true,
		},
		{
			name: "non-empty string with empty pattern",
			s:    "a",
			p:    "",
			want: false,
		},
		// {
		// 	name: "trailing Kleene star matching multiple characters",
		// 	s:    "abcddd",
		// 	p:    "abcd*",
		// 	want: true,
		// },
		// Simple cases with wildcard
		{name: "single wildcard match", s: "a", p: ".", want: true},
		{name: "single wildcard mismatch", s: "ab", p: ".", want: false},

		// Complex cases with wildcard
		{name: "wildcard match in middle", s: "abcde", p: "a.c.e", want: true},
		{name: "wildcard no match", s: "abc", p: "a.d", want: false},

		// Mixed wildcard and Kleene star
		{name: "wildcard with Kleene star, full match", s: "abcde", p: "a.*e", want: true},
		{name: "wildcard with Kleene star, partial match", s: "abcde", p: "a.*f", want: false},

		// Kleene star and wildcard only in pattern
		{name: "only Kleene and wildcard in pattern", s: "anything", p: ".*", want: true},
		{name: "complex wildcard and Kleene combination", s: "beginningmiddleend", p: "beg.*nd", want: true},

		// Edge cases
		{name: "pattern shorter than string", s: "abcd", p: "a.*f", want: false},
		{name: "string longer than pattern", s: "abc", p: "a.", want: false},
		{name: "no match due to Kleene end mismatch", s: "beginning", p: "beginning.*end", want: false},
		{name: "match with multiple wildcards and stars", s: "abxyzudee", p: "a.*x.z.d.*e", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
