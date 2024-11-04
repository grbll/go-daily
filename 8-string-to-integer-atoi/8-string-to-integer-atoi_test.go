package stringtointegeratoi

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "very large neg int", s: "-91283472332", want: -2147483648},
		{name: "positive number", s: "123", want: 123},
		{name: "negative number", s: "-123", want: -123},
		{name: "number with leading spaces", s: "   -42", want: -42},
		{name: "number with trailing characters", s: "4193 with words", want: 4193},
		{name: "overflow positive", s: "2147483648", want: int(maxInt)},
		{name: "overflow negative", s: "-2147483649", want: int(minInt)},
		{name: "leading zeros", s: "0000123", want: 123},
		{name: "empty string", s: "", want: 0},
		{name: "string with only spaces", s: "     ", want: 0},
		{name: "invalid characters", s: "abc123", want: 0},
		{name: "mixed invalid characters", s: "  -0012abc", want: -12},
		{name: "just a plus sign", s: "+", want: 0},
		{name: "just a minus sign", s: "-", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
