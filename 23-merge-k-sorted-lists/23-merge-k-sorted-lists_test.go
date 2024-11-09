package mergeksortedlists

import (
	"reflect"
	"testing"

	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_insert(t *testing.T) {
	type args struct {
		lists []*ListNode
		node  *ListNode
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			name: "Insert into empty list",
			args: args{
				lists: []*ListNode{},
				node:  BuildListNode([]int{5}),
			},
			want: []*ListNode{BuildListNode([]int{5})},
		},
		{
			name: "Insert at the beginning",
			args: args{
				lists: []*ListNode{BuildListNode([]int{10}), BuildListNode([]int{20})},
				node:  BuildListNode([]int{5}),
			},
			want: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10}), BuildListNode([]int{20})},
		},
		{
			name: "Insert in the middle",
			args: args{
				lists: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{20})},
				node:  BuildListNode([]int{10}),
			},
			want: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10}), BuildListNode([]int{20})},
		},
		{
			name: "Insert at the end",
			args: args{
				lists: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10})},
				node:  BuildListNode([]int{20}),
			},
			want: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10}), BuildListNode([]int{20})},
		},
		{
			name: "Insert duplicate value",
			args: args{
				lists: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10})},
				node:  BuildListNode([]int{10}),
			},
			want: []*ListNode{BuildListNode([]int{5}), BuildListNode([]int{10}), BuildListNode([]int{10})},
		},
		{
			name: "Insert multiple values into sorted list",
			args: args{
				lists: []*ListNode{BuildListNode([]int{1}), BuildListNode([]int{3}), BuildListNode([]int{5}), BuildListNode([]int{7})},
				node:  BuildListNode([]int{4}),
			},
			want: []*ListNode{BuildListNode([]int{1}), BuildListNode([]int{3}), BuildListNode([]int{4}), BuildListNode([]int{5}), BuildListNode([]int{7})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.lists, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pop(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Pop from single-element list",
			args: args{
				lists: []*ListNode{BuildListNode([]int{5})},
			},
			want: BuildListNode([]int{5}),
		},
		{
			name: "Pop from multi-element list",
			args: args{
				lists: []*ListNode{BuildListNode([]int{1}), BuildListNode([]int{3})},
			},
			want: BuildListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pop(&tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Merge single list",
			args: args{
				lists: []*ListNode{BuildListNode([]int{1, 2, 3})},
			},
			want: BuildListNode([]int{1, 2, 3}),
		},
		{
			name: "Merge multiple lists",
			args: args{
				lists: []*ListNode{
					BuildListNode([]int{1, 4, 5}),
					BuildListNode([]int{1, 3, 4}),
					BuildListNode([]int{2, 6}),
				},
			},
			want: BuildListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
