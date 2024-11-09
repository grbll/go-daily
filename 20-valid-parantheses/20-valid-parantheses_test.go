package validparantheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid simple parentheses",
			input: "()",
			want:  true,
		},
		{
			name:  "Valid multiple types",
			input: "()[]{}",
			want:  true,
		},
		{
			name:  "Valid nested",
			input: "({[]})",
			want:  true,
		},
		{
			name:  "Invalid closing without opening",
			input: "(]",
			want:  false,
		},
		{
			name:  "Invalid unbalanced parentheses",
			input: "([)]",
			want:  false,
		},
		{
			name:  "Single unclosed parenthesis",
			input: "(",
			want:  false,
		},
		{
			name:  "Empty string",
			input: "",
			want:  true,
		},
		{
			name:  "Unclosed nested brackets",
			input: "{[}",
			want:  false,
		},
		{
			name:  "Valid complex nested pattern",
			input: "{[()()](({}))}",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
