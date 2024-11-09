package removenthnodefromendoflist

import (
	. "github.com/grbll/go-daily/singly-linked-list"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "remove 1st from end of 5-node list",
			input: []int{1, 2, 3, 4, 5},
			n:     1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "remove 2nd from end of 5-node list",
			input: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "remove head node (5th from end)",
			input: []int{1, 2, 3, 4, 5},
			n:     5,
			want:  []int{2, 3, 4, 5},
		},
		{
			name:  "remove single element from 1-node list",
			input: []int{1},
			n:     1,
			want:  []int{},
		},
		{
			name:  "remove 1st from end of 2-node list",
			input: []int{1, 2},
			n:     1,
			want:  []int{1},
		},
		{
			name:  "remove 2nd from end of 2-node list (remove head)",
			input: []int{1, 2},
			n:     2,
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildListNode(tt.input)
			got := removeNthFromEnd(head, tt.n)
			gotSlice := BuildSlice(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", tt.input, tt.n, gotSlice, tt.want)
			}
		})
	}
}
