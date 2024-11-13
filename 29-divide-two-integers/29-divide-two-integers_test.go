package dividetwointegers

import "testing"

func Test_divide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{name: "Positive Division", dividend: 10, divisor: 2, want: 5},
		{name: "Negative Dividend", dividend: -10, divisor: 2, want: -5},
		{name: "Negative Divisor", dividend: 10, divisor: -2, want: -5},
		{name: "Both Negative", dividend: -10, divisor: -2, want: 5},
		{name: "Divide By One", dividend: 5, divisor: 1, want: 5},
		{name: "Divide By Itself", dividend: 5, divisor: 5, want: 1},
		{name: "Zero Dividend", dividend: 0, divisor: 3, want: 0},
		{name: "Large Dividend", dividend: 2147483647, divisor: 2, want: 1073741823},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.dividend, tt.divisor)
			if got != tt.want {
				t.Errorf("%s: divide(%d, %d) = %d, want %d", tt.name, tt.dividend, tt.divisor, got, tt.want)
			}
		})
	}
}
