package mergeintervals

import (
	"reflect"
	"testing"
)

// func Test_rise(t *testing.T) {
// 	type args struct {
// 		heap [][]int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rise(tt.args.heap)
// 		})
// 	}
// }
//
// func Test_sink(t *testing.T) {
// 	type args struct {
// 		heap [][]int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sink(tt.args.heap)
// 		})
// 	}
// }

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "all to one",
			args: args{intervals: [][]int{{2, 9}, {1, 10}, {4, 5}, {3, 8}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "all to two",
			args: args{intervals: [][]int{{14, 16}, {2, 9}, {1, 10}, {4, 5}, {3, 8}}},
			want: [][]int{{1, 10}, {14, 16}},
		},
		{
			name: "close merge",
			args: args{intervals: [][]int{{14, 15}, {14, 16}}},
			want: [][]int{{14, 16}},
		},
		{
			name: "all to two, close merge",
			args: args{intervals: [][]int{{14, 15}, {2, 9}, {1, 10}, {14, 16}, {4, 5}, {3, 8}}},
			want: [][]int{{1, 10}, {14, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
