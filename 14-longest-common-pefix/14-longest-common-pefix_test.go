package longestcommonpefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "common prefix", input: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "single character prefix", input: []string{"dog", "racecar", "car"}, want: ""},
		{name: "all identical strings", input: []string{"test", "test", "test"}, want: "test"},
		{name: "one string is prefix of others", input: []string{"pre", "prefix", "prelude"}, want: "pre"},
		{name: "all strings are empty", input: []string{"", "", ""}, want: ""},
		{name: "some empty strings", input: []string{"", "prefix", "prelude"}, want: ""},
		{name: "no common prefix", input: []string{"apple", "banana", "cherry"}, want: ""},
		{name: "single string in array", input: []string{"solo"}, want: "solo"},
		{name: "common prefix with varying lengths", input: []string{"interstellar", "internet", "internal"}, want: "inter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
