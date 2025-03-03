package partitionlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one switch",
			args: args{head: BuildListNode([]int{1, 3, 2}), x: 3},
			want: BuildListNode([]int{1, 2, 3}),
		},
		{
			name: "preorderedlist",
			args: args{head: BuildListNode([]int{1, 2, 3}), x: 3},
			want: BuildListNode([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !cmp.Equal(tt.want, got) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
