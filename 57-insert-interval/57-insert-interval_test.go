package insertinterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "empty intervals",
			args: args{intervals: [][]int{}, newInterval: []int{5, 7}},
			want: [][]int{{5, 7}},
		},
		{
			name: "direct append",
			args: args{intervals: [][]int{{3, 4}}, newInterval: []int{5, 7}},
			want: [][]int{{3, 4}, {5, 7}},
		},
		{
			name: "direct prepend",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{0, 0}},
			want: [][]int{{0, 0}, {1, 5}},
		},
		{
			name: "inclusion",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{2, 3}},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
