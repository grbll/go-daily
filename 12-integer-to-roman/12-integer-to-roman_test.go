package integertoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		// Basic cases
		{name: "1", num: 1, want: "I"},
		{name: "4", num: 4, want: "IV"},
		{name: "9", num: 9, want: "IX"},
		{name: "10", num: 10, want: "X"},
		{name: "40", num: 40, want: "XL"},
		{name: "50", num: 50, want: "L"},
		{name: "90", num: 90, want: "XC"},
		{name: "100", num: 100, want: "C"},
		{name: "400", num: 400, want: "CD"},
		{name: "500", num: 500, want: "D"},
		{name: "900", num: 900, want: "CM"},
		{name: "1000", num: 1000, want: "M"},

		// More complex cases
		{name: "58", num: 58, want: "LVIII"},
		{name: "1994", num: 1994, want: "MCMXCIV"},
		{name: "2023", num: 2023, want: "MMXXIII"},
		{name: "3999", num: 3999, want: "MMMCMXCIX"}, // Largest standard Roman numeral value

		// Edge cases
		{name: "0 (invalid case)", num: 0, want: ""},
		{name: "maximum 3999", num: 3999, want: "MMMCMXCIX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Errorf("intToRoman(%d) = %s, want %s", tt.num, got, tt.want)
			}
		})
	}
}
