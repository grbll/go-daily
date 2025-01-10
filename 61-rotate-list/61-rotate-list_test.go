package rotatelist

import (
	cmp "github.com/google/go-cmp/cmp"
	"testing"

	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one element, one shift",
			args: args{head: BuildListNode([]int{1}), k: 1},
			want: BuildListNode([]int{1}),
		},
		{
			name: "two, element, one shift",
			args: args{head: BuildListNode([]int{1, 2}), k: 1},
			want: BuildListNode([]int{2, 1}),
		},
		{
			name: "five by two",
			args: args{head: BuildListNode([]int{1, 2, 3, 4, 5}), k: 2},
			want: BuildListNode([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "five by five",
			args: args{head: BuildListNode([]int{1, 2, 3, 4, 5}), k: 5},
			want: BuildListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "five by seven",
			args: args{head: BuildListNode([]int{1, 2, 3, 4, 5}), k: 7},
			want: BuildListNode([]int{4, 5, 1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !cmp.Equal(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
