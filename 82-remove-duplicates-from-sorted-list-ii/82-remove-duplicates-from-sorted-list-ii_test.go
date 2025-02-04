package removeduplicatesfromsortedlistii

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "remove first",
			args: args{head: BuildListNode([]int{1, 1, 2, 3})},
			want: BuildListNode([]int{2, 3}),
		},
		{
			name: "remove second",
			args: args{head: BuildListNode([]int{1, 2, 2, 3})},
			want: BuildListNode([]int{1, 3}),
		},
		{
			name: "remove nothing",
			args: args{head: BuildListNode([]int{1, 2, 3})},
			want: BuildListNode([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !cmp.Equal(got, tt.want) {
				t.Errorf("deleteDuplicates(%v) = %v, want %v", BuildSlice(tt.args.head), got, tt.want)
			}
		})
	}
}
