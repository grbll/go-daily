package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "single symbol", input: "I", want: 1},
		{name: "simple addition", input: "III", want: 3},
		{name: "subtractive notation", input: "IV", want: 4},
		{name: "larger subtractive notation", input: "IX", want: 9},
		{name: "standard number", input: "LVIII", want: 58},    // L=50, V=5, III=3
		{name: "complex number", input: "MCMXCIV", want: 1994}, // M=1000, CM=900, XC=90, IV=4
		{name: "high value", input: "MMMCMXCIX", want: 3999},   // maximum typical Roman numeral: 1000 + 1000 + 1000 + 900 + 90 + 9
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			if got != tt.want {
				t.Errorf("romanToInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
