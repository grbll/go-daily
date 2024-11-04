package reverseinteger

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "positive number within range", x: 123, want: 321},
		{name: "negative number within range", x: -123, want: -321},
		{name: "positive number exceeding int32 limit", x: 1534236469, want: 0},
		{name: "negative number exceeding int32 limit", x: -1534236469, want: 0},
		{name: "number with trailing zeros", x: 120, want: 21},
		{name: "zero input", x: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
