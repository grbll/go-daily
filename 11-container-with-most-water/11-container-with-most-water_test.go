package containerwithmostwater

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "simple case", height: []int{1, 1}, want: 1},
		{name: "increasing heights", height: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "decreasing heights", height: []int{5, 4, 3, 2, 1}, want: 6},
		{name: "mixed heights", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "plateau", height: []int{3, 3, 3, 3}, want: 9},
		{name: "single element", height: []int{4}, want: 0}, // Not enough lines to form a container
		{name: "alternating heights", height: []int{2, 1, 2, 1, 2, 1, 2}, want: 12},
		{name: "large peak in middle", height: []int{1, 2, 4, 8, 4, 2, 1}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}
