package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantLen int
		wantArr []int
	}{
		{
			name:    "No duplicates",
			input:   []int{1, 2, 3, 4, 5},
			wantLen: 5,
			wantArr: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "All elements the same",
			input:   []int{1, 1, 1, 1},
			wantLen: 1,
			wantArr: []int{1},
		},
		{
			name:    "Mixed duplicates",
			input:   []int{1, 1, 2, 3, 3, 4},
			wantLen: 4,
			wantArr: []int{1, 2, 3, 4},
		},
		{
			name:    "Duplicates at end",
			input:   []int{1, 2, 3, 4, 4, 4},
			wantLen: 4,
			wantArr: []int{1, 2, 3, 4},
		},
		{
			name:    "Empty list",
			input:   []int{},
			wantLen: 0,
			wantArr: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input) // Preserve original input

			gotLen := removeDuplicates(inputCopy)
			gotArr := inputCopy[:gotLen]

			if gotLen != tt.wantLen {
				t.Errorf("removeDuplicates() length = %v, want %v", gotLen, tt.wantLen)
			}
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("removeDuplicates() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
